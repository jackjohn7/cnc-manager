package main

import (
	"go.uber.org/fx"

	"github.com/JingusJohn/kyle-cnc/backend/controllers"
	"github.com/JingusJohn/kyle-cnc/backend/server"
)

func main() {
	fx.New(
		fx.Provide(
			// create general handlers
			controllers.CreateGeneralController,
			controllers.CreateAuthController,
		),
		fx.Invoke(server.CreateServer),
	).Run()
}
