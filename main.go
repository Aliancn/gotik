package main

import (
	"gotik/api"
	"gotik/api/login"
	"gotik/api/register"
	"gotik/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 建立表
	err := api.DB.AutoMigrate(&model.Comment{})
	if err != nil {
		panic(err)
	}
	err = api.DB.AutoMigrate(&model.Message{})
	if err != nil {
		panic(err)
	}
	err = api.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	err = api.DB.AutoMigrate(&model.UserFavorite{})
	if err != nil {
		panic(err)
	}
	err = api.DB.AutoMigrate(&model.UserFollow{})
	if err != nil {
		panic(err)
	}
	err = api.DB.AutoMigrate(&model.UserFriends{})
	if err != nil {
		panic(err)
	}
	err = api.DB.AutoMigrate(&model.UserVideo{})
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.POST("/douyin/user/register/", register.RegisterHandler)
	engine.POST("/douyin/user/login/", login.LoginHandler)

	engine.Run("0.0.0.0:8888")
}
