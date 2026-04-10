package dao

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// Database is the database instance for DAO layer
	Database gdb.DB
)

// Init initializes the database connection using g.DB().
// It reads configuration from manifest/config/config.yaml.
func Init(ctx context.Context) error {
	// Use g.DB() to initialize the database connection
	// The configuration will be automatically loaded from manifest/config/config.yaml
	Database = g.DB()
	
	g.Log().Info(ctx, "database connection initialized successfully")
	return nil
}
