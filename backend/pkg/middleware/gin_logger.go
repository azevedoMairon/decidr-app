package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLoggerJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		slog.Info("request",
			"status", status,
			"method", method,
			"path", path,
			"query", raw,
			"ip", clientIP,
			"latency", latency.Seconds(),
			"error", errorMessage,
		)
	}
}
