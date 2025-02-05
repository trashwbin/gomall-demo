package dal

import (
	"github.com/trashwbin/gomall-demo/app/user/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
