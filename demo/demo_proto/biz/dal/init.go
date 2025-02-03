package dal

import (
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
