package providers

import (
	"drones/app"
	"fmt"
)

type Provider interface {
	Boot() error
}

func Ignite(a *app.App) {
	err := NewRouteProvider(a).Boot()
	if err != nil {
		panic(fmt.Sprintf("failed to ignite route provider; Err: %v", err))
	}
	err = NewSchedulerProvider(a).Boot()
	if err != nil {
		panic(fmt.Sprintf("failed to ignite scheduler provider; Err: %v", err))
	}
}
