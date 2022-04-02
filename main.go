package main

import (
	"drones/app"
	"drones/app/providers"
)

func main() {
	a := app.NewApp()

	providers.Ignite(a)

	a.Run()
}
