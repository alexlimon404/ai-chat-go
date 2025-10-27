package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Index GET /go-api/ping
func (h *PingHandler) Index(c *gin.Context) {
	startTime := time.Now()

	responseTime := time.Since(startTime).Nanoseconds()

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"responseTime": responseTime,
		},
	})
}
