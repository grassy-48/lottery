package point

import (
	"fmt"
	"lottery/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResCheckStore struct {
	UID      int    `json:"user_id"`
	Code     string `json:"code"`
	CanStore bool   `json:"canStore"`
}

func CheckStore(c echo.Context) error {
	ckey := c.FormValue("code")
	uid, _ := strconv.Atoi(c.Param("userID"))
	if uid == 0 {
		return fmt.Errorf("user not found")
	}
	canStore := models.CheckCanStore(uid, ckey)
	res := ResCheckStore{
		UID:      uid,
		Code:     ckey,
		CanStore: canStore,
	}
	return c.JSON(http.StatusOK, res)
}

func Store(c echo.Context) error {
	ckey := c.FormValue("code")
	var point models.LtPoint
	uid, _ := strconv.Atoi(c.Param("userID"))
	var user models.LtUser
	if uid == 0 {
		return fmt.Errorf("user not found")
	}
	user = models.LtUser{
		Model: models.Model{ID: uid},
	}

	// user-code check
	if !models.CheckCanStore(uid, ckey) {
		return fmt.Errorf("deplicate code")
	}

	// code -> point
	fmt.Println(uid, ckey, user)
	models.GetPointFromCode(ckey, &point)

	if point.Point <= 0 {
		return fmt.Errorf("invalid code")
	}

	// user ++ point
	models.StorePointToUser(point.Point, &user)
	// history
	models.InsertPlusPointHistory(&user, ckey, &point)
	return c.JSON(http.StatusOK, user)
}
