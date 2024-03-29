package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GeneralController struct{}

func CreateGeneralController() *GeneralController {
	return &GeneralController{}
}

func (controller *GeneralController) Register(app *echo.Echo) {
	app.GET("/health", healthcheck)
}

func healthcheck(c echo.Context) error {
	log.Println("Health checked!")
	return c.String(http.StatusOK, "Healthy!")
}
