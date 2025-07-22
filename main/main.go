package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/yusufguntav/crm-client/listener"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LanguageType string

const (
	Turkish LanguageType = "tr"
	English LanguageType = "en"
)

type User struct {
	gorm.Model
	Username               string       `json:"username" gorm:"not null"`
	PhoneNumber            string       `json:"phone_number" gorm:"not null;uniqueIndex"`
	PhoneNumberIsVerified  bool         `json:"phone_number_is_verified" gorm:"default:false"`
	CountryCode            string       `json:"country_code" gorm:"not null"`
	Password               string       `json:"password" gorm:"not null"`
	Mail                   string       `json:"mail" gorm:"not null;uniqueIndex"`
	MailIsVerified         bool         `json:"mail_is_verified" gorm:"default:false"`
	LastCustomerAssignDate sql.NullTime `json:"last_customer_assign_date" gorm:"default:null"`
	Income                 float64      `json:"income" gorm:"default:0"`
}

type UserDetail struct {
	gorm.Model
	UserID    uint         `json:"user_id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	BirthDate sql.NullTime `json:"birth_date"`
	Address   string       `json:"address"`
	City      string       `json:"city"`
	State     string       `json:"state"`
	Language  LanguageType `json:"language" gorm:"type:varchar(5);default:'tr'"`
}

func main() {
	dsn := "postgres://postgres:password@localhost:5455/testdb?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM bağlantı hatası:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("SQL DB hatası:", err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&User{}, &UserDetail{}); err != nil {
		log.Fatal("Migration hatası:", err)
	}

	log.Println("Veritabanı tabloları başarıyla oluşturuldu!")

	tableData := listener.TableData{
		"users": {
			"email": "email",
			"id":    "id_in_project",
			"name":  "name",
			"phone": "phone",
		},
		"user_details": {
			"user_id":   "id_in_project",
			"birthday":  "birthday",
			"wp_credit": "special_fields.wp_credit",
		},
	}

	listener.StartDatabaseListener(
		tableData,
		dsn,
		"http://localhost:8080/api/v1/customer/callback",
		"de677ca7-308e-4a4c-ae20-deb813e30158",
	)

	select {}

}
