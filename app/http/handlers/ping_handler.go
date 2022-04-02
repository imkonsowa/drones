package handlers

import "github.com/gin-gonic/gin"

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (p *PingHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
	return
}
