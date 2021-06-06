package models

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "lottery/config"
    "time"
)

type Model struct {
    ID        uint        `gorm:"primary_key" json:"id"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt *time.Time  `json:"updated_at"`
    DeletedAt *time.Time  `json:"deleted_at"`
}

type LtUser struct {
    Model
	Mail string `json:"mail"`
	Name string `json:"name"`
	Circle string `json:"circle"`
	IsParticipant bool `json:"is_participant"`
	IsCreator bool `json:"is_creator"`	
	Place string `json:"place"`
	Point int `json:"point"`	
	Total int `json:"total"`
	Minus int `json:"minus"`
}

var Db *gorm.DB

func GetAllUsers(users *[]LtUser) {
    Db.Find(&users)
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
