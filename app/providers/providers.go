package providers

import "drones/app"

type Provider interface {
	Boot() error
}

func Ignite(a *app.App) {
	NewRouteProvider(a).Boot()
}
