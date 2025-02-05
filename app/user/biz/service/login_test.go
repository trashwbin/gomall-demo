package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/trashwbin/gomall-demo/app/user/biz/dal/mysql"
	user "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "1demo@damin.com",
		Password: "FJODIAFUFJO",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
