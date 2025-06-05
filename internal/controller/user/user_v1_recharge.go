package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error) {
	user := service.Session().GetUser(ctx)
	if user == nil {
		return nil, gerror.New("unauthorized")
	}
	newBalance, err := service.User().Recharge(ctx, int(user.Id), req.Amount)
	if err != nil {
		return nil, err
	}
	return &v1.RechargeRes{NewBalance: newBalance}, nil
}
