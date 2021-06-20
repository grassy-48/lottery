package code

import (
	"lottery/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Check(c echo.Context) error {
	cstr := c.QueryParam("code")
	var code models.LtCode
	if cstr == "" {
		return c.JSON(http.StatusBadRequest, code)
	}
	code = models.LtCode{
		UniqKey: cstr,
	}
	models.GetCode(&code)
	if code.ID == 0 {
		return c.JSON(http.StatusNoContent, code)
	}
	return c.JSON(http.StatusOK, code)
}
