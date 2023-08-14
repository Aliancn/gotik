package main

import (
	"gotik/api/favorite"
	"gotik/api/getcommentlist"
	"gotik/api/getfavorlist"
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
	engine.GET("/douyin/favorite/list/", getfavorlist.GetFavorListHandler)
	engine.GET("/douyin/comment/list/", getcommentlist.GetFavorListHandler)
	engine.POST("/douyin/favorite/action/", favorite.FavoriteHandler)

	engine.Run("0.0.0.0:8888")
}
