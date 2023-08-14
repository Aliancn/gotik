package favorite

import (
	"github.com/gin-gonic/gin"
	"gotik/api"
	svc_favorite "gotik/service/favorite"
	err_comm "gotik/utils/error_codes/common"
	err_favor "gotik/utils/error_codes/favorite"
	"gotik/utils/token"
	"net/http"
	"strconv"
)

type Output struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

const (
	favor   = 1
	unFavor = 2
)

func FavoriteHandler(ctx *gin.Context) {
	output := Output{}

	// 获取视频ID
	videoIDStr := ctx.PostForm("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		panic(err)
	}

	// 用户鉴权
	var logined = false
	tk := ctx.Query("token")
	tkInfo, err := token.GetTokenInfoFromToken(tk)
	if err == nil {
		logined = true
	}

	// 未登录
	if !logined {
		output.StatusCode = err_comm.ErrCodeNotLogin
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeNotLogin)
		ctx.JSON(http.StatusOK, &output)
		return
	}

	// 获取userID和option
	userID := tkInfo.UserID
	optionStr := ctx.PostForm("action_type")
	option, err := strconv.Atoi(optionStr)
	if err != nil {
		panic(err)
	}

	// 根据ID进行操作
	switch option {
	case favor:
		resultDoFavor, err := svc_favorite.DOFavor(api.DB, userID, videoID)
		if err != nil {
			panic(err)
		}
		if resultDoFavor.Code == svc_favorite.ResultCodeOptionErr {
			output.StatusCode = err_favor.ErrCodeOptionWrong
			output.StatusMsg = err_favor.GetStatusMessage(err_favor.ErrCodeOptionWrong)
			ctx.JSON(http.StatusOK, &output)
			return
		}
	case unFavor:
		resultDoUnFavor, err := svc_favorite.DOUnFavor(api.DB, userID, videoID)
		if err != nil {
			panic(err)
		}
		if resultDoUnFavor.Code == svc_favorite.ResultCodeOptionErr {
			output.StatusCode = err_favor.ErrCodeOptionWrong
			output.StatusMsg = err_favor.GetStatusMessage(err_favor.ErrCodeOptionWrong)
			ctx.JSON(http.StatusOK, &output)
			return
		}
	default:
		output.StatusCode = err_favor.ErrCodeOptionTypeWrong
		output.StatusMsg = err_favor.GetStatusMessage(err_favor.ErrCodeOptionTypeWrong)
		ctx.JSON(http.StatusOK, &output)
		return
	}

}
