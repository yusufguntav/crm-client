package listener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

func StartDatabaseListener(tableData TableData, dsn, crmURL, projectSecret string) {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println("Listener hatası:", err)
		}
	}

	listener := pq.NewListener(dsn, 10, 30, reportProblem)

	// Her tablo için ayrı kanal dinle
	for tableName := range tableData {
		if err := listener.Listen(fmt.Sprintf("%s_changes", tableName)); err != nil {
			log.Fatalf("Listen hatası (%s): %v", tableName, err)
		}
	}

	go func() {
		for {
			select {
			case n := <-listener.Notify:
				if n != nil {
					var extra map[string]interface{}
					if err := json.Unmarshal([]byte(n.Extra), &extra); err != nil {
						log.Println("JSON parse hatası:", err)
						continue
					}
					handleDynamicNotify(extra, tableData, crmURL, projectSecret)
				}
			}
		}
	}()
}
func handleDynamicNotify(extra map[string]interface{}, tableData TableData, crmURL, secret string) {
	tableName, ok := extra["table"].(string)
	if !ok {
		log.Println("Tablo bilgisi yok")
		return
	}

	columns, tableExists := tableData[tableName]
	if !tableExists {
		log.Println("Desteklenmeyen tablo:", tableName)
		return
	}

	body := make(map[string]interface{})
	var specialFields []map[string]interface{}

	for column, targetField := range columns {
		value, exists := extra[column]
		if !exists {
			log.Printf("Kolon (%s) veride bulunamadı.", column)
			continue
		}

		valueStr := fmt.Sprintf("%v", value)

		if strings.HasPrefix(targetField.(string), "special_fields.") {
			fieldName := strings.TrimPrefix(targetField.(string), "special_fields.")
			specialFields = append(specialFields, map[string]interface{}{
				"field": fieldName,
				"value": valueStr,
			})
		} else {
			body[targetField.(string)] = valueStr
		}
	}

	if len(specialFields) > 0 {
		body["special_fields"] = specialFields
	}

	sendCallback(body, crmURL, secret)
}
func sendCallback(body map[string]interface{}, crmURL, secret string) {
	data, err := json.Marshal(body)
	if err != nil {
		log.Println("JSON marshal hatası:", err)
		return
	}

	// Gönderilecek JSON'u logla
	log.Printf("Giden callback isteği: URL=%s\nBody=%s", crmURL, string(data))

	req, err := http.NewRequest("POST", crmURL, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Callback isteği hazırlanamadı:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ProjectSecretKey", secret)

	// Header'ları logla
	for k, v := range req.Header {
		log.Printf("Header: %s = %s\n", k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Callback isteği başarısız:", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Callback sonucu: %v", resp.Status)
}
