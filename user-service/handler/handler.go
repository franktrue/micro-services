package handler

import (
	"errors"
	pb "github.com/franktrue/micro-services/user-service/proto/user"
	"github.com/franktrue/micro-services/user-service/repo"
	"github.com/franktrue/micro-services/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
)

type UserService struct {
	Repo  repo.Repository
	Token service.Authable
}

func NewUserService(r repo.Repository, t service.Authable) *UserService {
	return &UserService{Repo: r, Token: t}
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, user *pb.User, res *pb.Response) error {
	// 对密码进行哈希加密
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	if err := srv.Repo.Create(user); err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *UserService) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	// 获取用户信息
	user, err := srv.Repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	// 校验用户输入密码是否于数据库存储密码匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	// 生成 jwt token
	token, err := srv.Token.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	// 校验用户亲求中的token信息是否有效
	claims, err := srv.Token.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("无效的用户")
	}

	res.Valid = true

	return nil
}
