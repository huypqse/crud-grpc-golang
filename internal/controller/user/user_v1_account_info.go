package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) AccountInfo(ctx context.Context, req *v1.AccountInfoReq) (res *v1.AccountInfoRes, err error) {
	user := service.Session().GetUser(ctx)
	if user == nil {
		return nil, gerror.New("unauthorized")
	}
	balance, err := service.User().GetBalance(ctx, int(user.Id))
	if err != nil {
		return nil, err
	}
	return &v1.AccountInfoRes{Balance: balance}, nil
}
