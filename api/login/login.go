package login

import (
	"encoding/json"
	"gotik/api"
	svc_login "gotik/service/login"
	err_comm "gotik/utils/error_codes/common"
	err_login "gotik/utils/error_codes/login"
	verifyinput "gotik/utils/verify_input"

	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginInputJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginOutputJSON struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     uint   `json:"user_id"`
	Token      string `json:"token"`
}

func LoginHandler(ctx *gin.Context) {
	outputData := loginOutputJSON{}

	inputBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		outputData.StatusCode = err_comm.ErrCodeInvalidArgs
		outputData.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(http.StatusOK, &outputData)
		return
	}

	inputData := loginInputJSON{}
	if err := json.Unmarshal(inputBytes, &inputData); err != nil {
		outputData.StatusCode = err_comm.ErrCodeInvalidArgs
		outputData.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(http.StatusOK, &outputData)
		return
	}

	// 这里判断一下而不查数据库是为了减小数据库的压力, 把能判断的东西尽量判断了
	if !verifyinput.IsUsernameValid(inputData.Username) || !verifyinput.IsPasswordValid(inputData.Password) {
		outputData.StatusCode = err_comm.ErrCodeInvalidArgs
		outputData.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(http.StatusOK, &outputData)
		return
	}

	result, err := svc_login.DoLogin(api.DB, inputData.Username, inputData.Password)
	if err != nil {
		panic(err)
	}

	switch result.Code {
	case svc_login.ResultPasswordWrong, svc_login.ResultUserNotExists:
		outputData.StatusCode = err_login.ErrCodeUsernameOrPasswordWrong
		outputData.StatusMsg = err_login.GetStatusMessage(err_login.ErrCodeUsernameOrPasswordWrong)

		ctx.JSON(200, &outputData)
		return
	case svc_login.ResultOK:
		outputData.StatusCode = err_comm.ErrCodeOK
		outputData.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
		outputData.Token = result.Token
		outputData.UserID = result.UserID

		ctx.JSON(200, &outputData)
		return
	default:
		panic("unexpected")
	}
}
