package dao

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// Database is the database instance for DAO layer
	Database gdb.DB
)

// Init initializes the database connection
func Init() error {
	Database = g.DB()
	return nil
}
