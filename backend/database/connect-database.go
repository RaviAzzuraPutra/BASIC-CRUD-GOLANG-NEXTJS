package database

import (
	"backend/config/db_config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var errorConnect error

	postgresDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		db_config.ConfigDB().Host, db_config.ConfigDB().User, db_config.ConfigDB().Password, db_config.ConfigDB().Name, db_config.ConfigDB().Port, db_config.ConfigDB().SSLMode, db_config.ConfigDB().Timezone,
	)

	DB, errorConnect = gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})

	if errorConnect != nil {
		panic("Terjadi Error Di connect-database.go: " + errorConnect.Error())
	}

	if DB == nil {
		log.Panic("Database Tidak Terhubung")
	}

	log.Printf("Berhasil Terhubung Ke Database ✅")
}
