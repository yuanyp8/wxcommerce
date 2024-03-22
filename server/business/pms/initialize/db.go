package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/business/pms/config"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
	"log"
	"os"
	"time"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 18:52
 * @Desc:
 */

// var DB *gorm.DB

func InitDB() {
	var (
		err error
		c   = config.C().MysqlInfo
	)
	dsn := fmt.Sprintf(constants.MySqlDSN, c.User, c.Password, c.Host, c.Port, c.Database)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true, // 设置为true表示GORM将预编译SQL语句，这有助于提高查询性能和防止SQL注入攻击
			SkipDefaultTransaction: true, // 设置为true意味着GORM不会自动为每个操作包裹一个默认的事务。开发者需要手动管理事务
			NamingStrategy: schema.NamingStrategy{ // 用于配置表名的命名规则，其中SingularTable: true表示在生成表名时，即使模型名是复数形式，也将映射到单数形式的数据库表名
				SingularTable: true,
			},
			Logger: newLogger,
		},
	)
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}

	if err = db.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
	}
}
