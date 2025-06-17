package controllers

import (
	"net/http"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/services"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/providers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AuthController struct {
	ControllerShared
	AccountService *services.AccountService
	jwt            *providers.JWTProvider
}

func NewAuthController(name string, accountService *services.AccountService, jwt *providers.JWTProvider) *AuthController {
	return &AuthController{
		ControllerShared: ControllerShared{Name: name},
		AccountService:   accountService,
		jwt:              jwt,
	}
}

type JSONObject map[string]interface{}

func (me *AuthController) RegisterAccount(c echo.Context) error {
	log.Info("Starting Signup")
	var info services.RegisterInfo
	err := c.Bind(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	created, err := me.AccountService.SignUp(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	if !created {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": "could not sign up user",
		})
	}
	return c.JSON(http.StatusOK, JSONObject{
		"message": "Account created successfully, check your email to activate",
	})
}

type ActivationInfo struct {
	Token string `json:"token"`
}

func (me *AuthController) ActivateAccount(c echo.Context) error {
	var info ActivationInfo
	err := c.Bind(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	err = me.AccountService.Activate(info.Token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, JSONObject{
		"message": "Account activated successfully, you can login now",
	})
}

func (me *AuthController) Login(c echo.Context) error {
	var info services.LoginInfo
	err := c.Bind(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	account, err := me.AccountService.Login(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	token, err := me.jwt.CreateJWT(account)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	log.Infof("%v user logged in successfully", info.Email)
	return c.JSON(http.StatusOK, JSONObject{
		"token": token,
	})
}

func (me *AuthController) ResetPassword(c echo.Context) error {
	var info services.ResetPasswordInfo
	err := c.Bind(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}
	err = me.AccountService.ResetPassword(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONObject{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, JSONObject{
		"message": "Password reset successfully, you can login now.",
	})
}

func (me *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/signup", me.RegisterAccount)
	e.POST("/activate", me.ActivateAccount)
	e.POST("/login", me.Login)
	e.POST("/resetpassword", me.ResetPassword)
}
