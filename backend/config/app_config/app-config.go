package app_config

import (
	"log"
	"os"
)

var PORT string

func ConfigAPP() {
	PORT = os.Getenv("APP_PORT")

	if PORT == "" {
		log.Println("APP_PORT tidak ditemukan")
	}
}
