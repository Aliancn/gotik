package register

import (
	"gotik/api"
	svcregister "gotik/service/register"
	err_comm "gotik/utils/error_codes/common"
	err_register "gotik/utils/error_codes/register"
	"gotik/utils/token"
	verifyinput "gotik/utils/verify_input"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerOutputJSON struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id"`
	Token      string `json:"token"`
}

func RegisterHandler(ctx *gin.Context) {
	// 用于作为返回的数据
	outputData := registerOutputJSON{}

	username := ctx.Query("username")
	password := ctx.Query("password")

	// 检查字段的格式
	if !verifyinput.IsUsernameValid(username) || !verifyinput.IsPasswordValid(password) {
		outputData.StatusCode = err_comm.ErrCodeInvalidArgs
		outputData.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(http.StatusOK, &outputData)
		return
	}

	// 开始进行业务
	result, err := svcregister.DoRegister(api.DB, username, password)

	// mysql在本地, 不可能是网络错误, 我们无能为力直接退出
	if err != nil {
		panic(err)
	}

	switch result.Code {
	case svcregister.ResultOK:
		outputData.StatusCode = err_comm.ErrCodeOK
		outputData.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		outputData.UserID = result.ID
		outputData.Token = token.NewToken(result.ID, username)

		ctx.JSON(200, &outputData)
		return
	case svcregister.ResultUsernameOccupied:
		outputData.StatusCode = err_register.ErrCodeUsernameOccupied
		outputData.StatusMsg = err_register.GetStatusMessage(err_register.ErrCodeUsernameOccupied)

		ctx.JSON(200, &outputData)
		return
	// 下面是防御式编程, 防止代码跑到本不应该到达的地方
	default:
		panic("unreachable")
	}
}
