package dal

import (
	"github.com/trashwbin/gomall-demo/app/frontend/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
