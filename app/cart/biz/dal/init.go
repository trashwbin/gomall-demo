package dal

import (
	"github.com/trashwbin/gomall-demo/app/cart/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
