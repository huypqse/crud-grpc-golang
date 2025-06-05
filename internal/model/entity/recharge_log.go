// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeLog is the golang structure for table recharge_log.
type RechargeLog struct {
	Id        int         `json:"id"        orm:"id"         description:"Recharge ID"`
	UserId    uint        `json:"userId"    orm:"user_id"    description:"User ID"`
	Amount    float64     `json:"amount"    orm:"amount"     description:"Recharge amount"`
	Before    float64     `json:"before"    orm:"before"     description:"Balance before recharge"`
	After     float64     `json:"after"     orm:"after"      description:"Balance after recharge"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	Remark    string      `json:"remark"    orm:"remark"     description:"Remark"`
}
