package cors_middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	// Baca daftar origin dari environment ALLOWED_ORIGINS dipisah koma tanpa spasi
	originsString := os.Getenv("ALLOWED_ORIGINS")
	if originsString == "" {
		// fallback default (development)
		originsString = "http://localhost:3000"
	}

	var allowOrigins []string

	if originsString != "" {
		allowOrigins = strings.Split(originsString, ",")
	}

	log.Println("Origin:", originsString)
	log.Println("Allowed origins:", allowOrigins)

	return func(ctx *gin.Context) {
		isOriginAllowed := func(origin string, allowOrigin []string) bool {
			for _, alloallowOrigin := range allowOrigin {
				if origin == alloallowOrigin {
					return true
				}
			}
			return false
		}

		origin := ctx.Request.Header.Get("Origin")

		if isOriginAllowed(origin, allowOrigins) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			log.Println("Header CORS dikirim utk", origin)
			ctx.Writer.Header().Set("Vary", "Origin")
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

			if ctx.Request.Method == http.MethodOptions {
				ctx.AbortWithStatus(http.StatusNoContent)
				return
			}
		}

		ctx.Next()
	}

}
