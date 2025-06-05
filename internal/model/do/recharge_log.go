// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeLog is the golang structure of table recharge_log for DAO operations like Where/Data.
type RechargeLog struct {
	g.Meta    `orm:"table:recharge_log, do:true"`
	Id        interface{} // Recharge ID
	UserId    interface{} // User ID
	Amount    interface{} // Recharge amount
	Before    interface{} // Balance before recharge
	After     interface{} // Balance after recharge
	CreatedAt *gtime.Time //
	Remark    interface{} // Remark
}
