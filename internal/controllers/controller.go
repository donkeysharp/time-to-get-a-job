package controllers

import (
	"net/http"
	"strconv"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControllerShared struct {
	Name     string
	Settings *web.Settings
}

func (me *ControllerShared) GetName() string {
	return me.Name
}

func (me *ControllerShared) SetSettings(settings *web.Settings) {
	me.Settings = settings
}

const INVALID_USER = -255

func GetUserId(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	subject, err := user.Claims.GetSubject()
	if err != nil {
		log.Errorf("Failed to retrieve jwt subject")
		return INVALID_USER, c.JSON(http.StatusInternalServerError, JSONObject{
			"message": "Internal error",
		})
	}

	userId, err := strconv.Atoi(subject)
	if err != nil {
		return INVALID_USER, c.JSON(http.StatusBadRequest, JSONObject{
			"message": "Invalid jwt",
		})
	}
	return userId, nil
}
