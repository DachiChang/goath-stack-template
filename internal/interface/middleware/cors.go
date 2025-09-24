package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	}
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}

	return cors.New(corsConfig)
}
