package providers

import (
	"drones/app"
	"drones/app/http/handlers"
	"drones/app/http/router"
)

type RouteProvider struct {
	App *app.App
}

func NewRouteProvider(a *app.App) *RouteProvider {
	return &RouteProvider{
		App: a,
	}
}

func (r *RouteProvider) Boot() error {
	pingHandler := handlers.NewPingHandler()
	dronesHandler := handlers.NewDronesHandler()

	rtr := &router.Router{
		App: r.App,

		PingHandler:   pingHandler,
		DronesHandler: dronesHandler,
	}

	rtr.RegisterAPIRoutes()

	return nil
}
