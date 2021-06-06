package controllers

import (
	"github.com/labstack/echo/v4"
    "lottery/models"
    "net/http"
)

func GetAllUsers(c echo.Context) error {
    var items []models.LtUser
    models.GetAllUsers(&items)
	return c.JSON(http.StatusOK, items)
}


