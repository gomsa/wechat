package hander

import (
	"context"

	pb "github.com/gomsa/wechat/proto/wechat"
)

// Wechat 授权服务处理
type Wechat struct {
	AppId  string
	Secret string
}

// ProcessCommonRequest 处理公共请求
func (srv *Wechat) ProcessCommonRequest(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {

	return nil
}
