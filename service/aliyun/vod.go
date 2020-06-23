package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"time"
)

func InitVodClient() (client *vod.Client, err error) {
	// 点播服务接入区域
	regionId := "cn-shanghai"
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     "",
		AccessKeySecret: "",
	}

	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = false          // 失败是否自动重试
	config.MaxRetryTime = 3          // 最大重试次数
	config.Timeout = time.Second * 3 // 连接超时，单位：纳秒；默认为3秒

	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

func MyGetPlayInfo(client *vod.Client, videoId string) (response *vod.GetPlayInfoResponse, err error) {
	// 创建API请求并设置参数，调用vod.Create${apiName}Request
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	// 发起请求并处理异常，调用client.${apiName}(request)
	return client.GetPlayInfo(request)
}
