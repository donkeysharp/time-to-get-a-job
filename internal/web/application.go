package web

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Settings struct {
	Port        int
	BindAddress string
}

type WebApplication struct {
	e        *echo.Echo
	Settings *Settings
}

type Controller interface {
	RegisterRoutes(e *echo.Echo)
	GetName() string
}

func NewWebApplication(settings *Settings) *WebApplication {
	e := echo.New()

	return &WebApplication{
		e:        e,
		Settings: settings,
	}
}

func (app *WebApplication) RegisterController(controller Controller) {
	log.Infof("Registering controller: %v", controller.GetName())
	controller.RegisterRoutes(app.e)
}

func (app *WebApplication) Start() {
	log.Info("Starting application")
	listenAddress := fmt.Sprintf("%v:%v", app.Settings.BindAddress, app.Settings.Port)
	app.e.Logger.Fatal(app.e.Start(listenAddress))
}
