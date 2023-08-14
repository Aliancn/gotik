package getfavorlist

import (
	"github.com/gin-gonic/gin"
	"gotik/api"
	svc_getfavorlist "gotik/service/getfavorlist"
	svc_getuserinfo "gotik/service/getuserinfo"
	svc_getworknumber "gotik/service/getworknumber"
	err_comm "gotik/utils/error_codes/common"
	"gotik/utils/token"
	"net/http"
	"strconv"
)

type getFavorListOutput struct {
	StatusCode int     `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户点赞视频列表
}

// Video
type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// 视频作者信息
// User
type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

func GetFavorListHandler(ctx *gin.Context) {
	output := getFavorListOutput{}

	// 获取用户ID
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		output.VideoList = nil
		ctx.JSON(http.StatusOK, &output)
	}

	// 用户鉴权
	var logined = false
	tk := ctx.Query("token")
	tkInfo, err := token.GetTokenInfoFromToken(tk)
	if err == nil {
		logined = true
	}
	if !logined {
		output.StatusCode = err_comm.ErrCodeNotLogin
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeNotLogin)
		output.VideoList = nil
		ctx.JSON(http.StatusOK, &output)
	}

	// 根据ID查询Favor列表
	resultFav, err := svc_getfavorlist.DoGetFavorList(api.DB, userID)
	if err != nil {
		panic(err)
	}

	if resultFav.Code == svc_getfavorlist.ResultCodeTargetNotFound {
		output.StatusCode = err_comm.ErrCodeOK
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		output.VideoList = nil
		ctx.JSON(200, &output)
		return
	}
	videoList := resultFav.VideoList

	// 遍历每个video
	// 查询对应的用户信息
	var result []Video
	for _, video := range videoList {
		authorID := video.AuthorID
		resultAut, err := svc_getuserinfo.DoGetUserInfo(api.DB, int(tkInfo.UserID), int(authorID))
		if err != nil {
			panic(err)
		}

		if resultAut.Code == svc_getuserinfo.ResultCodeTargetNotFound {
			output.StatusCode = err_comm.ErrCodeOK
			output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
			output.VideoList = nil
			ctx.JSON(200, &output)
			return
		}
		authorInfo := resultAut.UserInfo

		var isFollow bool
		if int(authorInfo.ID) == tkInfo.UserID {
			isFollow = false
		} else {
			isFollow = resultAut.IsFollow
		}

		resultWork, err := svc_getworknumber.DOGetWorkNumber(api.DB, authorInfo.ID)
		if resultWork.Code == svc_getworknumber.ResultCodeTargetNotFound {
			resultWork.WorkCount = 0
		}
		workCount := resultWork.WorkCount

		result = append(result, Video{
			Author: User{
				Avatar:          authorInfo.AvatarURL,
				BackgroundImage: authorInfo.BackgroundImageURL,
				FavoriteCount:   int64(authorInfo.FavoriteCount),
				FollowCount:     int64(authorInfo.FollowCount),
				FollowerCount:   int64(authorInfo.FollowerCount),
				ID:              int64(authorInfo.ID),
				IsFollow:        isFollow,
				Name:            authorInfo.Username,
				Signature:       authorInfo.Signature,
				TotalFavorited:  strconv.Itoa(int(authorInfo.FavoritedCount)),
				WorkCount:       int64(workCount),
			},
			CommentCount:  int64(video.CommentCount),
			CoverURL:      video.CoverURL,
			FavoriteCount: int64(video.FavoriteCount),
			ID:            int64(video.ID),
			IsFavorite:    true, // 表结构没有该字段 但由于这里查询的是点赞的video 所以是true
			PlayURL:       video.PlayURL,
			Title:         video.Title,
		})
	}

	output.StatusCode = err_comm.ErrCodeOK
	output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
	output.VideoList = result
	ctx.JSON(200, &output)
}
