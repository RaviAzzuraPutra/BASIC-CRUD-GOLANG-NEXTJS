package bootstrap

import (
	"backend/config"
	"backend/config/app_config"
	"backend/database"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitialApp() {
	errENV := godotenv.Load()

	if errENV != nil {
		log.Println("Terjadi Kesalahan Saat Load .env")
	}

	config.IndexConfig()

	database.ConnectDatabase()

	app := gin.Default()

	routes.TravelRouter(app)

	app.Run(app_config.PORT)
}
