package handlers

import (
	"drones/app/data/adapters"
	"drones/app/data/models"
	"drones/app/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DronesHandler struct {
	DronesAdapter adapters.DronesAdapter
}

func NewDronesHandler(adapter adapters.DronesAdapter) *DronesHandler {
	return &DronesHandler{
		adapter,
	}
}

func (d *DronesHandler) Register(context *gin.Context) {
	ctxKey, _ := context.Get("request")
	request := ctxKey.(*requests.RegisterDrone)

	drone := &models.Drone{
		Model:           request.Model,
		SerialNumber:    request.SerialNumber,
		WeightLimit:     request.WeightLimit,
		Status:          request.Status,
		BatteryCapacity: request.BatteryCapacity,
	}

	drone = d.DronesAdapter.Create(drone)

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Drone registered successfully",
		"drone":   drone,
	})
}
