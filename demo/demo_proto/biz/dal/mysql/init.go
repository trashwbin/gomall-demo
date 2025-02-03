package mysql

import (
	"fmt"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/model"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 从环境变量中获取数据库连接信息
	dsn := fmt.Sprintf(
		conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}

	var version Version
	err = DB.Raw("SELECT VERSION() as version").Scan(&version).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql version: ", version)

	// 自动迁移数据库表结构
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("auto migrate user table failed: ", err)
		panic(err)
	}
}
