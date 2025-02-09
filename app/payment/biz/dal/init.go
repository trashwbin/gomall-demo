package dal

import (
	"github.com/trashwbin/gomall-demo/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
