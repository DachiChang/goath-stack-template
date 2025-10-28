package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthorizationMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
