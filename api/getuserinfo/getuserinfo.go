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

	tk := ctx.Query("token")
	tkInfo, err := token.GetTokenInfoFromToken(tk)
	if err == nil {
		tkInfo.UserID = -1
	}

	result, err := svc_getuserinfo.DoGetUserInfo(api.DB, tkInfo.UserID, userID)
	if err != nil {
		panic(err)
	}

	if result.Code == svc_getuserinfo.ResultCodeTargetNotFound {
		output.StatusCode = err_comm.ErrCodeOK
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		output.User = nil
		ctx.JSON(200, &output)
		return
	}
	userInfo := result.UserInfo

	output.StatusCode = err_comm.ErrCodeOK
	output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
	output.User = &userInfoJSON{
		ID:              userInfo.ID,
		Name:            userInfo.Username,
		FollowCount:     int(userInfo.FollowCount),
		FollowerCount:   int(userInfo.FollowerCount),
		IsFollow:        result.IsFollow,
		Avatar:          userInfo.AvatarURL,
		BackgroundImage: userInfo.AvatarURL,
		Signature:       userInfo.Signature,
		TotalFavorited:  int(userInfo.FavoritedCount),
		WorkCount:       int(userInfo.WorkCount),
		FavoriteCount:   int(userInfo.FavoriteCount),
	}
	ctx.JSON(200, &output)
}
