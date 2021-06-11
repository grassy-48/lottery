package models

import (
	"fmt"
	"log"
	"lottery/config"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type LtUser struct {
	Model
	Mail          string `json:"mail"`
	Name          string `json:"name"`
	Circle        string `json:"circle"`
	IsParticipant bool   `json:"is_participant"`
	IsCreator     bool   `json:"is_creator"`
	Place         string `json:"place"`
	Point         int    `json:"point"`
	Total         int    `json:"total"`
	Minus         int    `json:"minus"`
}

type LtGift struct {
	Model
	Name    string `json:"name"`
	Grade   string `json:"grade"`
	Booth   string `json:"booth"`
	IsStock bool   `json:"is_stock"`
}

type LtPoint struct {
	Model
	Type  string `json:"type"`
	Point int    `json:"point"`
}

type LtCode struct {
	Model
	UniqKey string `json:"uniq_key"`
	Path    string `json:"path"`
	Owner   int    `json:"owner"`
}

type LtCodesPoint struct {
	Model
	CodeID  int `json:"code_id"`
	PointID int `json:"point_id"`
}

type LtPointHistory struct {
	Model
	UserID int `json:"user_id"`
	CodeID int `json:"code_id"`
	Plus   int `json:"point"`
	Minus  int `json:"minus"`
	Result int `json:"result"`
}

type LtGiftHistory struct {
	Model
	UserID int `json:"user_id"`
	GiftID int `json:"gift_id"`
}

var Db *gorm.DB

func GetAllUsers(users *[]LtUser) {
	Db.Find(&users)
}
func GetUser(user *LtUser) {
	Db.Where(user).Or(LtUser{Mail: user.Mail}).Take(&user)
}
func UpsertUser(user *LtUser) *gorm.DB {
	var c int
	Db.Table("lt_users").Where(LtUser{Mail: user.Mail}).Count(&c)
	if c > 0 {
		Db.Model(&user).Updates(user)
	} else {
		Db.Where(LtUser{Mail: user.Mail}).Assign(user).FirstOrCreate(&user)
	}
	return Db.Find(&user)
}

func GetCodes(uid int, codes *[]LtCode) *gorm.DB {
	return Db.Find(&codes, LtCode{Owner: uid})
}

func FindInsertCodes(uid int, codes *[]LtCode) *gorm.DB {
	var c int
	Db.Table("lt_codes").Where(LtCode{Owner: uid}).Count(&c)
	if c > 0 {
		return Db.Find(&codes, LtCode{Owner: uid})
	} else {
		rs := []string{
			srand(16),
			srand(16),
			srand(16),
			srand(16),
			srand(16),
			srand(16),
		}

		for i := range rs {
			c := LtCode{
				Owner:   uid,
				UniqKey: rs[i],
				Path:    "/store?code=" + rs[i],
			}
			Db.Create(&c)
			pid := 1
			if i == (len(rs) - 1) {
				pid = 2
			}
			cp := LtCodesPoint{
				CodeID:  c.ID,
				PointID: pid,
			}
			Db.Create(&cp)
		}
	}
	return Db.Find(&codes, LtCode{Owner: uid})
}

func GetPointFromCode(code string, point *LtPoint) *gorm.DB {
	return Db.Table("lt_points").
		Select("lt_points.*").
		Joins("join lt_codes_points on lt_points.id = lt_codes_points.point_id").
		Joins("join lt_codes on lt_codes_points.code_id = lt_codes.id").
		Where("lt_codes.uniq_key = ?", code).
		Find(&point)
}

func StorePointToUser(point int, user *LtUser) {
	Db.First(&user)
	user.Point += point
	user.Total += point
	Db.Save(user)
	Db.Find(&user)
}

func UsePointToUser(point int, user *LtUser) {
	Db.First(&user)
	user.Point -= point
	user.Total -= point
	Db.Save(user)
	Db.Find(&user)
}

func InsertPlusPointHistory(user *LtUser, ckey string, point *LtPoint) error {
	var code LtCode
	Db.Find(&code, LtCode{UniqKey: ckey})
	if code.ID == 0 {
		return fmt.Errorf("invalid code key")
	}
	h := LtPointHistory{
		UserID: user.ID,
		CodeID: code.ID,
		Plus:   point.Point,
		Result: user.Point,
	}
	Db.Create(&h)
	return nil
}

func InsertMinusPointHistory(user *LtUser, point int) error {
	h := LtPointHistory{
		UserID: user.ID,
		Minus:  point,
		Result: user.Point,
	}
	Db.Create(&h)
	return nil
}

func CheckCanDraw(user *LtUser, dp int) (bool, error) {
	Db.Where(user).Take(&user)
	cp := false
	if user.Point >= dp {
		cp = true
	}
	return cp, nil
}

func UpdateToFalseGiftStatus(gift *LtGift) {
	Db.First(&gift)
	gift.IsStock = false
	Db.Save(&gift)
	Db.Find(&gift)
}

func InsertGiftHistory(user *LtUser, gift *LtGift) error {
	h := LtGiftHistory{
		UserID: user.ID,
		GiftID: gift.ID,
	}
	Db.Create(&h)
	return nil
}

func GetValidGiftIds() ([]int, error) {
	var gifts []LtGift
	Db.Where("is_stock = ?", true).
		Where("deleted_at is null").
		Select("id").
		Find(&gifts)

	var ids []int
	for _, g := range gifts {
		ids = append(ids, g.ID)
	}
	return ids, nil
}

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func srand(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

// データベースの初期化
func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)

	// configから読み込んだ情報を元に、データベースに接続します
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}
}
