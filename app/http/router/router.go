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
		drones.POST("/register", r.inject(requests.InjectableRegisterDroneRequest), r.DronesHandler.Register)
	}

	return r
}
