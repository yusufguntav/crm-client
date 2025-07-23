package crmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func (c *Client) StartDatabaseListener(tableData TableData, dsn, crmURL, projectSecret string) {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println("Listener hatası:", err)
		}
	}

	listener := pq.NewListener(dsn, 10, 30, reportProblem)
	crmURL = c.BaseURL + crmURL

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

	req, err := http.NewRequest("POST", crmURL, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Callback isteği hazırlanamadı:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ProjectSecretKey", secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Callback isteği başarısız:", err)
		return
	}
	defer resp.Body.Close()
}

func SetupTriggersFromTableData(db *gorm.DB, tableData TableData) error {
	for tableName, columns := range tableData {
		triggerFuncSQL := fmt.Sprintf(`
		CREATE OR REPLACE FUNCTION notify_%s_change()
		RETURNS TRIGGER AS $$
		DECLARE
			payload JSON;
		BEGIN
			-- DELETE işlemi
			IF TG_OP = 'DELETE' THEN
				payload := json_build_object(%s);
				PERFORM pg_notify('%s_changes', payload::text);
				RETURN OLD;
			END IF;

			-- INSERT veya UPDATE işlemi
			IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
				payload := json_build_object(%s);
				PERFORM pg_notify('%s_changes', payload::text);
				RETURN NEW;
			END IF;

			RETURN NULL;
		END;
		$$ LANGUAGE plpgsql;
		`,
			tableName,
			buildPayload(columns, "OLD"),
			tableName,
			buildPayload(columns, "NEW"),
			tableName,
		)

		if err := db.Exec(triggerFuncSQL).Error; err != nil {
			return fmt.Errorf("trigger function error on table %s: %w", tableName, err)
		}

		triggerSQL := fmt.Sprintf(`
		DROP TRIGGER IF EXISTS %s_change_trigger ON %s;
		CREATE TRIGGER %s_change_trigger
		AFTER INSERT OR UPDATE OR DELETE ON %s
		FOR EACH ROW EXECUTE FUNCTION notify_%s_change();
		`,
			tableName, tableName,
			tableName,
			tableName, tableName,
		)

		if err := db.Exec(triggerSQL).Error; err != nil {
			return fmt.Errorf("trigger creation error on table %s: %w", tableName, err)
		}
	}
	return nil
}

func buildPayload(columns map[string]any, scope string) string {
	var parts []string
	parts = append(parts, `'table', TG_TABLE_NAME`, `'operation', TG_OP`)

	for colName := range columns {
		parts = append(parts, fmt.Sprintf("'%s', %s.%s", colName, scope, colName))
	}

	return strings.Join(parts, ",\n")
}
