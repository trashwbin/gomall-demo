package dal

import (
	"github.com/trashwbin/gomall-demo/app/checkout/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
