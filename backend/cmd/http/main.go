package main

import (
	"log"

	"github.com/realfabecker/buck3t/internal/adapters"
	cordom "github.com/realfabecker/buck3t/internal/core/domain"
	corpts "github.com/realfabecker/buck3t/internal/core/ports"
)

func main() {
	if err := container.Container.Invoke(func(
		app corpts.HttpHandler,
		walletConfig *cordom.Config,
	) error {
		if err := app.Register(); err != nil {
			return err
		}
		return app.Listen(walletConfig.AppPort)
	}); err != nil {
		log.Fatalln(err)
	}
}
