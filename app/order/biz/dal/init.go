package dal

import (
	"github.com/trashwbin/gomall-demo/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
