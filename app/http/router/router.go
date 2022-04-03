package router

import (
	"drones/app"
	"drones/app/http/handlers"
	"drones/app/http/requests"
)

type Router struct {
	App *app.App

	PingHandler   *handlers.PingHandler
	DronesHandler *handlers.DronesHandler
}

func (r *Router) RegisterAPIRoutes() *Router {
	engine := r.App.Engine

	engine.GET("/ping", r.PingHandler.Ping)

	drones := engine.Group("/drones")
	{
		drones.POST("/register", r.inject(&requests.RegisterDrone{}), r.DronesHandler.Register)
		drones.POST("/load", r.inject(&requests.LoadMedications{}), r.DronesHandler.LoadMedications)
		// drones.PUT("/update-status", r.inject(&requests.UpdateDroneStatus{}), r.DronesHandler.UpdateStatus)
		drones.GET("/:serialNumber/battery", r.DronesHandler.GetBatteryCapacity)
	}

	return r
}
