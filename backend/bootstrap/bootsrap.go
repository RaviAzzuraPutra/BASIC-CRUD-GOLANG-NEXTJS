package bootstrap

import (
	"backend/config"
	"backend/config/app_config"
	"backend/database"
	"backend/middleware/cors_middleware"
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

	// Gunakan middleware CORS kustom yang membaca daftar origin dari env ALLOWED_ORIGINS
	app.Use(cors_middleware.CorsMiddleware())

	// (Opsional) Endpoint wildcard OPTIONS agar preflight selalu mendapat 204 + header
	app.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	routes.TravelRouter(app)

	app.Run(app_config.PORT)
}
