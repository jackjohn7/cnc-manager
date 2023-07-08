package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/JingusJohn/kyle-cnc/backend/controllers"
)

func CreateServer(
	lc fx.Lifecycle,
	generalController *controllers.GeneralController,
	authController *controllers.AuthController,
) *echo.Echo {
	app := echo.New()

	// In the future, these Register functions may take a Group
	//  instead of the entire server instance. This will allow
	//  for group-wide middleware application

	// Add Routes
	//  Root
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Register Controllers

	//  General
	generalController.Register(app)
	//  Auth
	authController.Register(app)

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			fmt.Println("Welcome to the backend!")
			go app.Start(":1323")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// TODO: close any database connections

			// TODO: close any gRPC or ws connections

			// close server
			return app.Shutdown(ctx)
		},
	})
	return app
}
