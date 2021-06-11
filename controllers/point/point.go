package point

import (
	"fmt"
	"lottery/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
