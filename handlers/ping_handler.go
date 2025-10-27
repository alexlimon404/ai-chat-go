package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Index GET /go-api/ping
func (h *PingHandler) Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": `pong`,
	})
}
