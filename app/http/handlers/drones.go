package handlers

import (
	"drones/app/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DronesHandler struct {
}

func NewDronesHandler() *DronesHandler {
	return &DronesHandler{}
}

func (d *DronesHandler) Register(context *gin.Context) {
	ctxKey, _ := context.Get("request")
	request := ctxKey.(*requests.RegisterDrone)

	// TODO: register the drone to the DB

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Drone registered successfully",
		"drone":   request,
	})
}
