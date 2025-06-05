package v1

import "github.com/gogf/gf/v2/frame/g"

type AccountInfoReq struct {
	g.Meta `path:"/account/info" method:"get" tags:"AccountService" summary:"Get account balance"`
}

type AccountInfoRes struct {
	Balance float64 `json:"balance"`
}
