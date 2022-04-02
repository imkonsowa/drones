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
	var request requests.RegisterDrone

	bindErr := context.Bind(&request)
	if bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"message": "Failed to parse your request",
		})
		return
	}

	// TODO: add validation

	// TODO: register the drone to the DB

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Drone registered successfully",
		"drone":   request,
	})
}
