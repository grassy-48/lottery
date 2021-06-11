package raffle

import (
	"fmt"
	"lottery/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
)

var drawPoint = 5

type resDraw struct {
	User    *models.LtUser `json:"user"`
	CanDraw bool           `json:"canDraw"`
}

type resElect struct {
	User *models.LtUser `json:"user"`
	Gift *models.LtGift `json:"gift"`
}

func DrawEvt(c echo.Context) error {
	uid, _ := strconv.Atoi(c.Param("userID"))
	var user models.LtUser
	if uid == 0 {
		return fmt.Errorf("user not found")
	}
	user = models.LtUser{
		Model: models.Model{ID: uid},
	}

	// user -- point
	models.UsePointToUser(drawPoint, &user)
	// history
	models.InsertMinusPointHistory(&user, drawPoint)
	return c.JSON(http.StatusOK, user)
}

func DrawOnl(c echo.Context) error {
	uid, _ := strconv.Atoi(c.Param("userID"))
	var user models.LtUser
	if uid == 0 {
		return fmt.Errorf("user not found")
	}
	user = models.LtUser{
		Model: models.Model{ID: uid},
	}

	// ポイント確認
	cp, err := models.CheckCanDraw(&user, drawPoint)
	if err != nil {
		return err
	}
	if !cp {
		return fmt.Errorf("lack of point")
	}

	// draw gift
	// 有効なgiftレコードを取得（idの配列）
	ids, err := models.GetValidGiftIds()
	pp.Println(ids)
	if err != nil {
		return err
	}
	// 乱数で抽選
	cid := ids[rand.Intn(len(ids))]
	pp.Println(cid)
	// 対象giftの在庫をfalseにしてgift情報取得
	gift := models.LtGift{
		Model: models.Model{ID: cid},
	}
	models.UpdateToFalseGiftStatus(&gift)
	// lt_gift_history 更新
	models.InsertGiftHistory(&user, &gift)

	// user -- point
	models.UsePointToUser(drawPoint, &user)
	// history
	models.InsertMinusPointHistory(&user, drawPoint)

	res := resElect{
		User: &user,
		Gift: &gift,
	}
	return c.JSON(http.StatusOK, res)
}

func Check(c echo.Context) error {
	uid, _ := strconv.Atoi(c.Param("userID"))
	var user models.LtUser
	if uid == 0 {
		return fmt.Errorf("user not found")
	}
	user = models.LtUser{
		Model: models.Model{ID: uid},
	}

	cp, err := models.CheckCanDraw(&user, drawPoint)
	if err != nil {
		return err
	}
	res := resDraw{
		User:    &user,
		CanDraw: cp,
	}

	return c.JSON(http.StatusOK, res)
}
