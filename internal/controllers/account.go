package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountController struct {
	ControllerShared
}

func NewAccountController(name string) *AccountController {
	return &AccountController{
		ControllerShared: ControllerShared{Name: name},
	}
}

func (me *AccountController) Get(c echo.Context) error {
	return c.String(http.StatusOK, "GetAccount")
}

func (me *AccountController) Update(c echo.Context) error {
	return c.String(http.StatusOK, "UpdateAccount")
}

func (me *AccountController) RegisterRoutes(e *echo.Echo) {
	e.GET("/account/me", me.Get)
	e.PUT("/account/me", me.Update)
}
