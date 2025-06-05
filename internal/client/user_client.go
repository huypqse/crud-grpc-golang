package service

import (
	"context"
	"io"

	"github.com/gogf/gf-demo-user/v2/backend-onboarding/api/proto"

	"google.golang.org/grpc"
)

type UserClient struct {
	client proto.UserServiceClient
}

// Tạo một client mới kết nối đến UserService
func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: proto.NewUserServiceClient(conn),
	}
}

// Gọi RPC GetUserById để lấy thông tin người dùng theo ID
func (c *UserClient) GetUserById(ctx context.Context, userId string) (*proto.UserResponse, error) {
	return c.client.GetUserById(ctx, &proto.GetUserRequest{UserId: userId})
}

// Gọi RPC GetUsersStream để nhận stream các thông tin người dùng từ danh sách ID
func (c *UserClient) GetUsersStream(ctx context.Context, userIds []string, streamChan chan<- *proto.UserResponse) error {
	stream, err := c.client.GetUsersStream(ctx, &proto.UserIdsRequest{UserIds: userIds})
	if err != nil {
		return err
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		streamChan <- user
	}

	return nil
}
