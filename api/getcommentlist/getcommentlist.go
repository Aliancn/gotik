package getcommentlist

import (
	"github.com/gin-gonic/gin"
	"gotik/api"
	svc_getcommentlist "gotik/service/getcommentlist"
	svc_getuserinfo "gotik/service/getuserinfo"
	svc_getworknumber "gotik/service/getworknumber"
	err_comm "gotik/utils/error_codes/common"
	"gotik/utils/token"
	"strconv"
)

type GetCommentListOutput struct {
	CommentList []Comment `json:"comment_list"` // 评论列表
	StatusCode  int64     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`   // 返回状态描述
}

// Comment
type Comment struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
}

// 评论用户信息
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
	output := GetCommentListOutput{}

	// 获取videoID
	videoIDStr := ctx.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		output.CommentList = nil
		ctx.JSON(200, &output)
	}

	// todo 用户鉴权
	var logined = false
	tk := ctx.Query("token")
	tkInfo, err := token.GetTokenInfoFromToken(tk)
	if err == nil {
		logined = true
	}
	// todo 对于没有登陆的怎么处理
	//if !logined {
	//
	//}

	// 根据ID查询comment列表
	resultCom, err := svc_getcommentlist.DoGetFavorList(api.DB, videoID)
	if err != nil {
		panic(err)
	}

	if resultCom.Code == svc_getcommentlist.ResultCodeTargetNotFound {
		output.StatusCode = err_comm.ErrCodeOK
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		output.CommentList = nil
		ctx.JSON(200, &output)
		return
	}
	commentList := resultCom.CommentList

	// 遍历每个video
	// 查询对应的用户信息
	var result []Comment
	for _, comment := range commentList {
		authorID := comment.UserID
		resultAut, err := svc_getuserinfo.DoGetUserInfo(api.DB, int(tkInfo.UserID), int(authorID))
		if err != nil {
			panic(err)
		}

		if resultAut.Code == svc_getuserinfo.ResultCodeTargetNotFound {
			output.StatusCode = err_comm.ErrCodeOK
			output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
			output.CommentList = nil
			ctx.JSON(200, &output)
			return
		}
		authorInfo := resultAut.UserInfo

		var isFollow bool
		if authorInfo.Username == tkInfo.Username || !logined {
			isFollow = false
		} else {
			isFollow = resultAut.IsFollow
		}

		resultWork, err := svc_getworknumber.DOGetWorkNumber(api.DB, authorInfo.ID)
		if resultWork.Code == svc_getworknumber.ResultCodeTargetNotFound {
			resultWork.WorkCount = 0
		}
		workCount := resultWork.WorkCount

		result = append(result, Comment{
			Content:    comment.Content,
			CreateDate: "",
			ID:         int64(comment.ID),
			User: User{
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
		})
	}

	output.StatusCode = err_comm.ErrCodeOK
	output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
	output.CommentList = result
	ctx.JSON(200, &output)
}
