package main

import (
	pb "github.com/franktrue/micro-services/user-service/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"golang.org/x/net/context"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your Email",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Your Password",
			},
		),
	)
	// 远程服务客户端调用句柄
	client := pb.NewUserService("micro-services.user.service", service.Client())

	// 运行客户端命令调用远程服务逻辑设置
	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")

			log.Println("参数", name, email, password)

			// 调用用户服务
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
			})

			if err != nil {
				log.Fatalf("创建用户失败：%v", err)
				return err
			}

			log.Printf("创建用户成功：%s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("获取所有用户失败：%v", err)
				return err
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatalf("用户客户端启动失败：%v", err)
	}
}
