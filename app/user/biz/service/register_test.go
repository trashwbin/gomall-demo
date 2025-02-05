package service

import (
	"context"
	user "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	//godotenv.Load("../../.env")
	//在这里调用mysql.Init() 会无法找到conf/test/conf.yaml,因为windows下在这里调用会导致路径不对，解决方式暂时是临时修改conf.go中的路径
	//mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "11demo@damin.com",
		Password:        "FJODIAFUFJO",
		PasswordConfirm: "FJODIAFUFJO",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
