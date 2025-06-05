package service

import (
	"context"
	"io"
	"log"

	"github.com/gogf/gf-demo-user/v2/backend-onboarding/api/proto"
	"google.golang.org/grpc"
)

type RechargeClient struct {
	client proto.RechargeLogServiceClient
}

// Khởi tạo RechargeClient với connection gRPC
func NewRechargeClient(conn *grpc.ClientConn) *RechargeClient {
	return &RechargeClient{
		client: proto.NewRechargeLogServiceClient(conn),
	}
}

// Gọi RPC CreateRechargeLog - gửi 1 request và nhận 1 response
func (c *RechargeClient) CreateRechargeLog(ctx context.Context, userId string, amount float64, remark string) (*proto.RechargeLogResponse, error) {
	return c.client.CreateRechargeLog(ctx, &proto.CreateRechargeRequest{
		UserId: userId,
		Amount: amount,
		Remark: remark,
	})
}

// Gọi RPC StreamRechargeLogs - gửi nhiều request và nhận nhiều response qua BiDi stream
func (c *RechargeClient) StreamRechargeLogs(ctx context.Context, requests []*proto.CreateRechargeRequest, streamChan chan<- *proto.RechargeLogResponse) error {
	stream, err := c.client.StreamRechargeLogs(ctx)
	if err != nil {
		return err
	}

	// Gửi stream trong 1 goroutine
	go func() {
		defer func() {
			_ = stream.CloseSend()
		}()
		for _, req := range requests {
			if err := stream.Send(req); err != nil {
				log.Printf("Send error: %v", err)
				return
			}
		}
	}()

	// Nhận stream liên tục
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		streamChan <- resp
	}

	return nil
}
