package getuserinfo

import (
	"gotik/api"
	"gotik/utils/token"
	"strconv"

	svc_getuserinfo "gotik/service/getuserinfo"
	err_comm "gotik/utils/error_codes/common"

	"github.com/gin-gonic/gin"
)

type userInfoJSON struct {
	ID              uint   `json:"id"`
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

type getUserInfoOutput struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`

	User *userInfoJSON `json:"user"`
}

func GetUserInfoHandler(ctx *gin.Context) {
	output := getUserInfoOutput{}

	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		output.User = nil
		ctx.JSON(200, &output)
		return
	}

	var logined = false
	var loginedUsername = ""
	tk := ctx.Query("token")
	tkInfo, err := token.GetTokenInfoFromToken(tk)
	if err == nil {
		logined = true
		loginedUsername = tkInfo.Username
	}

	user, err := svc_getuserinfo.DoGetUserInfo(api.DB, uint(userID))
	if err != nil {
		panic(err)
	}

	if user == nil {
		output.StatusCode = err_comm.ErrCodeOK
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		output.User = nil
		ctx.JSON(200, &output)
		return
	}

	if !logined || (logined && user.Username == loginedUsername) {
		output.StatusCode = err_comm.ErrCodeOK
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		output.User = &userInfoJSON{
			ID:              user.ID,
			Name:            user.Username,
			FollowCount:     int(user.FollowCount),
			FollowerCount:   int(user.FollowerCount),
			IsFollow:        false,
			Avatar:          user.AvatarURL,
			BackgroundImage: user.AvatarURL,
			Signature:       user.Signature,
			TotalFavorited:  int(user.FavoritedCount),
			WorkCount:       int(user.WorkCount),
			FavoriteCount:   int(user.FavoriteCount),
		}
		ctx.JSON(200, &output)
		return
	}

	isFollow, err := svc_getuserinfo.IsAFollowingB(api.DB, tkInfo.UserID, user.ID)
	if err != nil {
		panic(err)
	}

	output.StatusCode = err_comm.ErrCodeOK
	output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
	output.User = &userInfoJSON{
		ID:              user.ID,
		Name:            user.Username,
		FollowCount:     int(user.FollowCount),
		FollowerCount:   int(user.FollowerCount),
		IsFollow:        isFollow,
		Avatar:          user.AvatarURL,
		BackgroundImage: user.AvatarURL,
		Signature:       user.Signature,
		TotalFavorited:  int(user.FavoritedCount),
		WorkCount:       int(user.WorkCount),
		FavoriteCount:   int(user.FavoriteCount),
	}
	ctx.JSON(200, &output)
}
