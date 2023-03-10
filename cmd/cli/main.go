package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"recomends/internal/app"
	"recomends/internal/config"
)

func main() {
	log.Print("[Main] Startup, load config")
	cfg := config.GetConfig()

	kernel := app.NewKernel(cfg)
	kernel.
		ConfigureDatabase().
		ConfigureServiceLocator().
		AfterInitializationEvents().
		ConfigureCli()

	cliApp := kernel.DI.Get("cli").(*cli.App)

	kernel.ResolveContainer()

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
