package controllers

import (
	"errors"
	"net/http"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/services"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AccountController struct {
	ControllerShared
	AccountService *services.AccountService
}

func NewAccountController(name string, accountService *services.AccountService, settings *web.Settings) *AccountController {
	return &AccountController{
		ControllerShared: ControllerShared{Name: name, Settings: settings},
		AccountService:   accountService,
	}
}

func (me *AccountController) Get(c echo.Context) error {
	userId, err := GetUserId(c)
	if err != nil {
		return err
	}
	account, err := me.AccountService.GetProfile(userId)
	if err != nil {
		log.Errorf("Failed to retrieve profile information")
		return c.JSON(http.StatusInternalServerError, JSONObject{"message": "Internal server error"})
	}
	return c.JSON(http.StatusOK, account)
}

func (me *AccountController) Update(c echo.Context) error {
	userId, err := GetUserId(c)
	if err != nil {
		return err
	}
	var accountInfo services.AccountInfo
	if err := c.Bind(&accountInfo); err != nil {
		log.Warnf("Failed to bind accountinfo on update %v", err.Error())
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": "Bad request",
		})
	}
	if err := me.AccountService.UpdateProfile(&accountInfo, userId); err != nil {
		if errors.Is(err, services.ErrIncorrectFields) {
			return c.JSON(http.StatusBadRequest, JSONObject{
				"message": "Bad request",
			})
		}
		return c.JSON(http.StatusInternalServerError, JSONObject{
			"message": "Internal server error",
		})
	}
	return c.JSON(http.StatusOK, JSONObject{
		"message": "Account updated successfully",
	})
}

func (me *AccountController) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/account")

	settings := me.Settings
	g.Use(echojwt.JWT([]byte(settings.JWTSecret)))

	g.GET("/me", me.Get)
	g.PUT("/me", me.Update)
}
