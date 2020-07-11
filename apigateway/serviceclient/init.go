package serviceclient

import (
	"apigw/handler"
	proto "apigw/proto/account"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
)

var Port string

func RegisterService() {
	// 连接服务注册中心
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "p",
				Usage: "port",
			},
		),
	)
	// 解析命令行参数
	service.Init(
		micro.Action(func(c *cli.Context) error {
			Port = c.String("p")
			if len(Port) == 0 {
				Port = "8091"
			}
			return nil
		},
		),
	)
	// 复用服务注册的客户端
	//cli := client.DefaultClient
	// 获取在服务注册中心上 micro.service.account 的客户端
	handler.AccountServiceClient = proto.NewAccountService("micro.service.account", client.DefaultClient)
}
