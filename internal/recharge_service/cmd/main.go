package service

import (
	"context"
	"io"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gogf/gf-demo-user/v2/backend-onboarding/api/proto"
	"github.com/gogf/gf-demo-user/v2/backend-onboarding/internal/dao"
	"github.com/gogf/gf-demo-user/v2/backend-onboarding/internal/recharge_service/model"
)

type RechargeService struct {
	pb.UnimplementedRechargeLogServiceServer
	rechargeDao *dao.RechargeLogDao
	userClient  pb.UserServiceClient
}

func NewRechargeService(rechargeDao *dao.RechargeLogDao, userServiceAddr string) (*RechargeService, error) {
	conn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userClient := pb.NewUserServiceClient(conn)
	return &RechargeService{
		rechargeDao: rechargeDao,
		userClient:  userClient,
	}, nil
}

func (s *RechargeService) CreateRechargeLog(ctx context.Context, req *pb.CreateRechargeRequest) (*pb.RechargeLogResponse, error) {
	userResp, err := s.userClient.GetUserById(ctx, &pb.GetUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	recharge := &model.RechargeLog{
		UserID:    req.UserId,
		Amount:    req.Amount,
		Before:    0,
		After:     req.Amount,
		Remark:    req.Remark,
		CreatedAt: time.Now(),
	}

	id, err := s.rechargeDao.Create(recharge)
	if err != nil {
		return nil, err
	}

	return &pb.RechargeLogResponse{
		RechargeId:   strconv.Itoa(int(id)),
		UserId:       req.UserId,
		UserNickname: userResp.Nickname,
		Amount:       req.Amount,
		CreatedAt:    recharge.CreatedAt.Format(time.RFC3339),
		Remark:       req.Remark,
	}, nil
}

func (s *RechargeService) StreamRechargeLogs(stream pb.RechargeLogService_StreamRechargeLogsServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil // Kết thúc stream
		}
		if err != nil {
			return err
		}

		userResp, err := s.userClient.GetUserById(stream.Context(), &pb.GetUserRequest{UserId: req.UserId})
		if err != nil {
			return err
		}

		recharge := &model.RechargeLog{
			UserID:    req.UserId,
			Amount:    req.Amount,
			Before:    0,
			After:     req.Amount,
			Remark:    req.Remark,
			CreatedAt: time.Now(),
		}

		id, err := s.rechargeDao.Create(recharge)
		if err != nil {
			return err
		}

		resp := &pb.RechargeLogResponse{
			RechargeId:   strconv.Itoa(int(id)),
			UserId:       req.UserId,
			UserNickname: userResp.Nickname,
			Amount:       req.Amount,
			CreatedAt:    recharge.CreatedAt.Format(time.RFC3339),
			Remark:       req.Remark,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
