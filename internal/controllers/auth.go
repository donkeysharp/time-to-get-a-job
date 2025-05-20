package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	ControllerShared
}

func NewAuthController(name string) *AuthController {
	return &AuthController{
		ControllerShared: ControllerShared{Name: name},
	}
}

func (me *AuthController) RegisterAccount(c echo.Context) error {
	return c.String(http.StatusOK, "RegisterAccount")
}

func (me *AuthController) ActivateAccount(c echo.Context) error {
	return c.String(http.StatusOK, "ActivateAccount")
}

func (me *AuthController) Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

func (me *AuthController) ResetPassword(c echo.Context) error {
	return c.String(http.StatusOK, "ResetPassword")
}

func (me *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/signup", me.RegisterAccount)
	e.POST("/activate", me.ActivateAccount)
	e.POST("/login", me.Login)
	e.POST("/resetpassword", me.ResetPassword)
}
