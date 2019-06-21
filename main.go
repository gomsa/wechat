package main

import (
	// 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/tools/env"
	"github.com/gomsa/wechat/hander"
	pb "github.com/gomsa/wechat/proto/wechat"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 用户服务实现
	pb.RegisterWechatHandler(srv.Server(), &hander.Wechat{
		MiniprogramAppId:  env.Getenv("MINIPROGRAM_APP_ID", ""),
		MiniprogramSecret: env.Getenv("MINIPROGRAM_SECRET", ""),
	})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
