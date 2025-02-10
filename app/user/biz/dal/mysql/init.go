package mysql

import (
	"fmt"
	"github.com/trashwbin/gomall-demo/app/user/biz/model"
	"github.com/trashwbin/gomall-demo/app/user/conf"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	// 如果环境不是线上环境，则自动迁移数据库模式以匹配User模型
	// 这样做是为了在开发和测试环境中保持数据库结构的最新状态
	if conf.GetConf().Env != "online" {
		// 执行数据库模式自动迁移
		err = DB.AutoMigrate(&model.User{})
		// 如果迁移过程中发生错误，则抛出异常
		if err != nil {
			panic(err)
		}
	}
}
