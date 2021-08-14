package handler

import (
	pb "github.com/franktrue/micro-services/user-service/proto/user"
	"github.com/franktrue/micro-services/user-service/repo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type UserService struct {
	Repo repo.Repository
}

func NewUserService(r repo.Repository) *UserService {
	return &UserService{Repo: r}
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
