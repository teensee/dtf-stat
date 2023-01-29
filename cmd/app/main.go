package main

import (
	"log"
	"recomends/internal/app"
	"recomends/internal/config"
)

func main() {
	log.Print("Startup, load config")
	cfg := config.GetConfig()

	kernel := app.NewKernel(cfg)
	kernel.
		ConfigureDatabase().
		ConfigureServiceLocator().
		ConfigureRoutes().
		AfterInitializationEvents().
		ResolveContainer().
		End()
}
