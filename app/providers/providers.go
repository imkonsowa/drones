package providers

import "drones/app"

type Provider interface {
	Boot() error
}

func Ignite(a *app.App) {
	err := NewRouteProvider(a).Boot()
	if err != nil {
		panic("failed to ignite route provider")
	}
	err = NewSchedulerProvider(a).Boot()
	if err != nil {
		panic("failed to ignite scheduler provider")
	}
}
