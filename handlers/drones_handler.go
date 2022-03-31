package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterDrone(c *gin.Context) {
	var drone Drone
	err := c.Bind(drone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed to handle your request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Drone registered successfully!",
		"drone":   drone,
	})
}

func LoadDrone(c *gin.Context) {

}

func GetAvailableDrones(c *gin.Context) {

}

func GetDroneItems(c *gin.Context) {

}

func CheckDroneBatteryLevel(c *gin.Context) {

}

func UpdateDroneStatus(c *gin.Context) {

}
