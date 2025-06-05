package service

import (
	"context"

	pb "github.com/gogf/gf-demo-user/v2/backend-onboarding/api/proto"
	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf/v2/errors/gerror"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// Triển khai RPC GetUserById
func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// Sử dụng dao.User để truy vấn thông tin người dùng
	user, err := dao.User.Ctx(ctx).Where("id", req.UserId).One()
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.Newf(`User with id "%s" not found`, req.UserId)
	}

	return &pb.UserResponse{
		UserId:   user["id"].String(),
		Nickname: user["nickname"].String(),
	}, nil
}
