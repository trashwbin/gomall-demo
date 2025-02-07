package dal

import (
	"github.com/trashwbin/gomall-demo/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
