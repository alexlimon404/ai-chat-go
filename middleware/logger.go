package middleware

import (
	"ai-chat-go/database"
	"ai-chat-go/models"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				// Restore the body for further reading
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		c.Next()

		duration := time.Since(startTime)
		responseTimeNs := duration.Nanoseconds()
		responseTimeMs := float64(duration.Microseconds()) / 1000.0

		errorMessage := ""
		if len(c.Errors) > 0 {
			errorMessage = c.Errors.String()
		}

		logEntry := models.GoLog{
			Method:         c.Request.Method,
			Path:           c.Request.URL.Path,
			StatusCode:     c.Writer.Status(),
			ResponseTimeNs: responseTimeNs,
			ResponseTimeMs: responseTimeMs,
			IPAddress:      c.ClientIP(),
			UserAgent:      c.Request.UserAgent(),
			RequestBody:    requestBody,
			ResponseBody:   writer.body.String(),
			ErrorMessage:   errorMessage,
		}

		go func() {
			if err := database.DB.Create(&logEntry).Error; err != nil {
				println("Failed to save log entry:", err.Error())
			}
		}()
	}
}
