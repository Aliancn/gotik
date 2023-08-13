package api

import (
	"gotik/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:mark2004@tcp(127.0.0.1:8888)/gotik?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}

	DB = db

	// 建立表
	err = DB.AutoMigrate(&model.Comment{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.Message{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.UserFavorite{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.UserFollow{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.UserFriends{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.UserVideo{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.Video{})
	if err != nil {
		panic(err)
	}
}
