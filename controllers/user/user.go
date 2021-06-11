package user

import (
	"lottery/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResUser struct {
	User  *models.LtUser  `json:"user"`
	Codes []models.LtCode `json:"codes"`
}

func Create(c echo.Context) error {
	ut := c.FormValue("type")
	var place string
	var is_c bool
	var is_p bool
	if ut == "creator" {
		place = c.FormValue("place")
		is_c = true
	}
	if ut == "participant" {
		is_p = true
	}
	user := models.LtUser{
		Mail:          c.FormValue("mail"),
		Name:          c.FormValue("name"),
		Circle:        c.FormValue("circleName"),
		Place:         place,
		IsParticipant: is_p,
		IsCreator:     is_c,
	}
	models.UpsertUser(&user)
	if !user.IsCreator {
		return c.JSON(http.StatusOK, user)
	}
	// QRcode url path
	var codes []models.LtCode
	models.FindInsertCodes(int(user.Model.ID), &codes)
	res := ResUser{
		User:  &user,
		Codes: codes,
	}
	return c.JSON(http.StatusOK, res)
}

func Get(c echo.Context) error {
	uid, _ := strconv.Atoi(c.QueryParam("user_id"))
	var user models.LtUser
	if uid != 0 {
		user = models.LtUser{
			Model: models.Model{ID: uid},
		}
	} else if c.QueryParam("mail") != "" {
		user = models.LtUser{
			Mail: c.QueryParam("mail"),
		}
	}
	models.GetUser(&user)

	var codes []models.LtCode
	models.GetCodes(int(user.Model.ID), &codes)
	res := ResUser{
		User:  &user,
		Codes: codes,
	}
	return c.JSON(http.StatusOK, res)
}
