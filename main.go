package main

import (
	feed "gotik/api/getfeed"
	"gotik/api/getuserinfo"
	"gotik/api/login"
	"gotik/api/publish"
	"gotik/api/register"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	engine.POST("/douyin/user/register/", register.RegisterHandler)
	engine.POST("/douyin/user/login/", login.LoginHandler)
	engine.GET("/douyin/user/", getuserinfo.GetUserInfoHandler)
	engine.POST("/douyin/publish/action/", publish.PublishHandler)
	engine.GET("/douyin/feed/", feed.FeedHandler)

	engine.Run("0.0.0.0:8888")
}
