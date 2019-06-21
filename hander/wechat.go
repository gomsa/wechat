package hander

import (
	"context"

	sdk "github.com/bigrocs/wechat"
	"github.com/bigrocs/wechat/requests"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/wechat/proto/wechat"
)

// Wechat 授权服务处理
type Wechat struct {
	MiniprogramAppId  string
	MiniprogramSecret string
}

// ProcessCommonRequest 处理公共请求
func (srv *Wechat) ProcessCommonRequest(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {

	// 创建连接
	client, err := sdk.NewClient()
	if err != nil {
		log.Log(err)
		return err
	}
	client.Credential.Miniprogram.AppId = srv.MiniprogramAppId
	client.Credential.Miniprogram.Secret = srv.MiniprogramSecret
	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = req.Domain
	request.ApiName = req.ApiName
	request.QueryParams = req.QueryParams
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Log(err)
		return err
	}
	res.Content = response.GetHttpContentString()
	return
}
