package dal

import (
	"github.com/trashwbin/gomall-demo/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
