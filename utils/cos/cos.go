package cos

import (
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var Client *cos.Client

func init() {
	bucketURL, err := url.Parse("https://gotik-1257452121.cos.ap-chongqing.myqcloud.com")
	if err != nil {
		panic(err)
	}
	Client = cos.NewClient(&cos.BaseURL{BucketURL: bucketURL}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKID60gJitpdEaibIlodEdnkT6Xrxyay72jM",
			SecretKey: "ThhV6TecOwIQJzReiYCcolV1INEsIl01",
		},
	})
}
