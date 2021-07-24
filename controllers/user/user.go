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

//TODO: デプロイ時に切り替え
var dirBase = "/virtual/ymtk/public_html/ymtk.xyz/lottery/img/qr_code/"

//var dirBase = "./img/qr_code/"

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
	if user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, user)
	}
	if !user.IsCreator {
		res := ResUser{
			User: &user,
		}
		return c.JSON(http.StatusOK, res)
	}
	// QRcode url path
	var codes []models.LtCode
	models.FindInsertCodes(int(user.Model.ID), dirBase, &codes)
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
	models.GetUserCodes(int(user.Model.ID), &codes)
	res := ResUser{
		User:  &user,
		Codes: codes,
	}
	return c.JSON(http.StatusOK, res)
}
