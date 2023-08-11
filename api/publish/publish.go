package publish

import (
	"bytes"
	"context"
	"fmt"
	cosutil "gotik/utils/cos"
	"gotik/utils/fileutil"
	"gotik/utils/token"
	verifyinput "gotik/utils/verify_input"
	"os"
	"sync/atomic"

	err_comm "gotik/utils/error_codes/common"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

// 为了保证每个文件的唯一性, 此处做一个计数器
var fileCounter atomic.Int64

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

	ext, ok := fileutil.GetVideoFileExt(file.Filename)
	if !ok {
		output.StatusCode = err_comm.ErrCodeInvalidArgs
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(200, &output)
		return
	}

	currentCount := fileCounter.Add(1)
	saveFName := fmt.Sprintf("%d.%s", currentCount, ext)
	coverFName := fmt.Sprintf("%d-cover.png", currentCount)

	err = ctx.SaveUploadedFile(file, "tmp/"+saveFName)
	if err != nil {
		output.StatusCode = err_comm.ErrcodeInternalError
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeInvalidArgs)
		ctx.JSON(200, &output)
		return
	}

	buf := bytes.NewBuffer(nil)
	err = ffmpeg_go.Input("tmp/"+saveFName).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", 0)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		panic(err)
	}

	err = imaging.Save(img, "tmp"+coverFName)
	if err != nil {
		panic(err)
	}

	_, _, err = cosutil.Client.Object.Upload(context.Background(), saveFName, "tmp/"+saveFName, &cos.MultiUploadOptions{
		OptIni: &cos.InitiateMultipartUploadOptions{ACLHeaderOptions: &cos.ACLHeaderOptions{XCosACL: "public-read"}},
	})
	if err != nil {
		output.StatusCode = err_comm.ErrcodeInternalError
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrcodeInternalError)
		ctx.JSON(200, &output)
	}

	r, _, err := cosutil.Client.Object.Upload(context.Background(), saveFName, "tmp/"+coverFName, &cos.MultiUploadOptions{
		OptIni: &cos.InitiateMultipartUploadOptions{ACLHeaderOptions: &cos.ACLHeaderOptions{XCosACL: "public-read"}},
	})
	if err != nil {
		output.StatusCode = err_comm.ErrcodeInternalError
		output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrcodeInternalError)
		ctx.JSON(200, &output)
	}

	println(r.Location)

	output.StatusCode = err_comm.ErrCodeOK
	output.StatusMsg = err_comm.GetStatusMessage(err_comm.ErrCodeOK)
	ctx.JSON(200, &output)
}
