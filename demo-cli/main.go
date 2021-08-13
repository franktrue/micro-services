package main

import (
	"context"
	pb "github.com/franktrue/micro-services/demo-service/proto/demo"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("micro-services.demo.cli"),
	)
	service.Init()

	client := pb.NewDemoService("micro-services.demo.service", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "Franktrue"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}
