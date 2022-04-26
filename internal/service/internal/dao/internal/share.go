// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShareDao is the data access object for table wsx_share.
type ShareDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns ShareColumns // columns contains all the column names of Table for convenient usage.
}

// ShareColumns defines and stores column names for table wsx_share.
type ShareColumns struct {
	Id          string // 分享ID
	UserId      string // 用户ID
	Name        string // 分享名
	UpdateTime  string // 更新时间
	CreateTime  string // 创建时间
	Code        string // 提取码
	Times       string // 下载次数
	NeverExpire string // 1:不过期；0:过期
	ExpireTime  string // 过期时间
	Type        string // 1:文件夹；0:文件
	FileId      string // 文件文件夹ID
}

//  shareColumns holds the columns for table wsx_share.
var shareColumns = ShareColumns{
	Id:          "id",
	UserId:      "user_id",
	Name:        "name",
	UpdateTime:  "update_time",
	CreateTime:  "create_time",
	Code:        "code",
	Times:       "times",
	NeverExpire: "never_expire",
	ExpireTime:  "expire_time",
	Type:        "type",
	FileId:      "file_id",
}

// NewShareDao creates and returns a new DAO object for table data access.
func NewShareDao() *ShareDao {
	return &ShareDao{
		group:   "default",
		table:   "wsx_share",
		columns: shareColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShareDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShareDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShareDao) Columns() ShareColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShareDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShareDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShareDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
