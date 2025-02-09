package dal

import (
	"github.com/trashwbin/gomall-demo/app/email/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
