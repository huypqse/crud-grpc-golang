package v1

import "github.com/gogf/gf/v2/frame/g"

type RechargeReq struct {
	g.Meta `path:"/account/recharge" method:"post" tags:"AccountService" summary:"Recharge account"`
	Amount float64 `json:"amount" v:"required|min:0.01#Amount is required and must be greater than 0"`
}

type RechargeRes struct {
	NewBalance float64 `json:"new_balance"`
}
