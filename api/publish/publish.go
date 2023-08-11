package publish

import (
	"context"
	"fmt"
	"gotik/utils/cos"
	"gotik/utils/token"
	verifyinput "gotik/utils/verify_input"

	err_comm "gotik/utils/error_codes/common"

	"github.com/gin-gonic/gin"
)

type publishOutput struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func PublishHandler(ctx *gin.Context) {
	output := publishOutput{}

	tk := ctx.Request.FormValue("token")
	_, err := token.GetTokenInfoFromToken(tk)
	if err != nil {
		output.StatusCode = err_comm.ErrcodePermissionDenied
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrcodePermissionDenied)
		ctx.JSON(200, &output)
		return
	}

	title := ctx.Request.FormValue("title")
	if !verifyinput.IsTitleValid(title) {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(200, &output)
		return
	}

	file, err := ctx.FormFile("data")
	if err != nil {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(200, &output)
		return
	}

	f, err := file.Open()
	if err != nil {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(200, &output)
		return
	}

	resp, err := cos.Client.Object.Put(context.Background(), "1", f, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
