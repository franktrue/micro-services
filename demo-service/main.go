package main

import (
	"context"
	pb "github.com/franktrue/micro-services/demo-service/proto/demo"
	"github.com/micro/go-micro/v2"
	"log"
)

type DemoServiceHandler struct{}

func (d *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "你好，" + req.Name
	return nil
}

func main() {
	// 注册服务名必须和demo.proto中的package声明一致
	service := micro.NewService(
		micro.Name("micro-services.demo.service"),
	)

	service.Init()
	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
