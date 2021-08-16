/*
 * @Description: 入口文件
 * @Author: franktrue 807615827@qq.com
 * @Date: 2021-08-14 11:25:56
 * @LastEditTime: 2021-08-16 10:41:15
 */
package main

import (
	"github.com/franktrue/micro-services/user-service/db"
	"github.com/franktrue/micro-services/user-service/handler"
	pb "github.com/franktrue/micro-services/user-service/proto/user"
	"github.com/franktrue/micro-services/user-service/repo"
	"github.com/franktrue/micro-services/user-service/service"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	db, err := db.CreateConnect()
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}

	// 每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&pb.User{})

	repository := repo.NewUserRepository(db)

	token := service.NewTokenService(repository)

	userService := handler.NewUserService(repository, token)

	srv := micro.NewService(
		micro.Name("micro-services.user.service"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), userService)

	// 启动服务
	if err := srv.Run(); err != nil {
		log.Fatalf("user-service 启动失败：%v", err)
	}

}
