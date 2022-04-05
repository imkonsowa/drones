package providers

import (
	"drones/app"
	"drones/app/data/adapters"
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
	dronesAdapter := adapters.NewDronesAdapter(r.App.DB)
	medicationsAdapter := adapters.NewMedicationsAdapter(r.App.DB)

	pingHandler := handlers.NewPingHandler()
	dronesHandler := handlers.NewDronesHandler(dronesAdapter, medicationsAdapter)
	uploadsHandler := handlers.NewUploadsHandler()

	rtr := &router.Router{
		App: r.App,

		PingHandler:    pingHandler,
		DronesHandler:  dronesHandler,
		UploadsHandler: uploadsHandler,
	}

	rtr.RegisterAPIRoutes()

	return nil
}
