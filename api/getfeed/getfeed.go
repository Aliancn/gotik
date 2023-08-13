package feed

import (
	"gotik/api"
	svc_getfeed "gotik/service/getfeed"
	err_comm "gotik/utils/error_codes/common"
	"gotik/utils/token"
	"strconv"

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
	output := outputJSON{}

	lastTimeStr := ctx.Query("last_time")

	lastTime, err := strconv.Atoi(lastTimeStr)
	if err != nil {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(200, &output)
		return
	}

	tk := ctx.Query("token")
	tkInfo, err := token.GetTokenInfoFromToken(tk)
	if err != nil {
		tkInfo.UserID = -1
	}

	result, err := svc_getfeed.DoGetFeed(api.DB, tkInfo.UserID, lastTime)
	if err != nil {
		panic(err)
	}

	lis := result.List
	for _, v := range lis {
		output.VideoList = append(output.VideoList, &videoJSON{ID: int(v.Video.ID), Author: &authorJSON{
			ID:              int(v.Author.ID),
			Name:            v.Author.Username,
			FollowCount:     int(v.Author.FollowCount),
			FollowerCount:   int(v.Author.FollowerCount),
			IsFollow:        v.IsFollowAuthor,
			Avatar:          v.Author.AvatarURL,
			BackgroundImage: v.Author.BackgroundImageURL,
			Signature:       v.Author.Signature,
			TotalFavorited:  int(v.Author.FavoriteCount),
			WorkCount:       int(v.Author.WorkCount),
			FavoriteCount:   int(v.Author.FavoriteCount),
		},
			PlayURL:       v.Video.PlayURL,
			CoverUrl:      v.Video.CoverURL,
			FavoriteCount: int(v.Video.FavoriteCount),
			CommentCount:  int(v.Video.CommentCount),
			IsFavorite:    v.IsFavoriteVideo,
			Title:         v.Video.Title})
	}

	ctx.JSON(200, &output)
}
