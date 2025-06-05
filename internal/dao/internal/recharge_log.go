// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RechargeLogDao is the data access object for the table recharge_log.
type RechargeLogDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns RechargeLogColumns // columns contains all the column names of Table for convenient usage.
}

// RechargeLogColumns defines and stores column names for the table recharge_log.
type RechargeLogColumns struct {
	Id        string // Recharge ID
	UserId    string // User ID
	Amount    string // Recharge amount
	Before    string // Balance before recharge
	After     string // Balance after recharge
	CreatedAt string //
	Remark    string // Remark
}

// rechargeLogColumns holds the columns for the table recharge_log.
var rechargeLogColumns = RechargeLogColumns{
	Id:        "id",
	UserId:    "user_id",
	Amount:    "amount",
	Before:    "before",
	After:     "after",
	CreatedAt: "created_at",
	Remark:    "remark",
}

// NewRechargeLogDao creates and returns a new DAO object for table data access.
func NewRechargeLogDao() *RechargeLogDao {
	return &RechargeLogDao{
		group:   "default",
		table:   "recharge_log",
		columns: rechargeLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RechargeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RechargeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RechargeLogDao) Columns() RechargeLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RechargeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RechargeLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RechargeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
