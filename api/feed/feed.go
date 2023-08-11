package feed

import (
	"encoding/json"
	"fmt"
	"gotik/utils/consts"

	"github.com/gin-gonic/gin"
)

type authorJSON struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int    `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}

type videoJSON struct {
	ID            int         `json:"id"`
	Author        *authorJSON `json:"author"`
	PlayURL       string      `json:"play_url"`
	CoverUrl      string      `json:"cover_url"`
	FavoriteCount int         `json:"favorite_count"`
	CommentCount  int         `json:"comment_count"`
	IsFavorite    bool        `json:"is_vafovire"`
	Title         string      `json:"title"`
}

type outputJSON struct {
	StatusCode int          `json:"status_code"`
	StatusMsg  string       `json:"status_msg"`
	NextTime   int          `json:"next_time"`
	VideoList  []*videoJSON `json:"video_list"`
}

func FeedHandler(ctx *gin.Context) {
	m := &outputJSON{
		StatusCode: 0,
		StatusMsg:  "OK",
		NextTime:   30000,
		VideoList: []*videoJSON{{
			ID: 1,
			Author: &authorJSON{
				ID:              1,
				Name:            "markity",
				FollowCount:     0,
				FollowerCount:   0,
				IsFollow:        false,
				Avatar:          consts.DefaultAvatar,
				BackgroundImage: consts.DefaultBackground,
				Signature:       "默认签名",
				TotalFavorited:  0,
				WorkCount:       0,
				FavoriteCount:   0,
			},
			PlayURL:       "https://gotik-1257452121.cos.ap-chongqing.myqcloud.com/2023-06-01%2020-04-47.mp4",
			CoverUrl:      consts.DefaultBackground,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         "测试",
		},
		},
	}

	fmt.Println(json.Marshal(m))
	ctx.JSON(200, m)
}
