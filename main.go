package main

import (
	"drones/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	drones := router.Group("/drones")
	{
		drones.POST("/register", handlers.RegisterDrone)
		drones.GET("/available", handlers.GetAvailableDrones)
		drones.POST("/load", handlers.LoadDrone)
		drones.GET("/items", handlers.GetDroneItems)
		drones.GET("/battery-level", handlers.CheckDroneBatteryLevel)
		drones.PUT("/status", handlers.UpdateDroneStatus)
	}

	err := router.Run(":6060")
	if err != nil {
		panic("failed to run the app server")
	}
}
