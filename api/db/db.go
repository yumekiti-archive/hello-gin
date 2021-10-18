package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBconnect() *gorm.DB {
	DBMS := "mysql"
    USER := "user"
    PASS := "password"
    PROTOCOL := "tcp(db:3306)"
    DBNAME := "database"
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    }
    return db
}

type User struct {
    gorm.Model
    NickName string `json:"nickName"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&User{})
    return db
}
