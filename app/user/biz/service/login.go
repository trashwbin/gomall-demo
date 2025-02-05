package service

import (
	"context"
	"errors"
	"github.com/trashwbin/gomall-demo/app/user/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/user/biz/model"
	user "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

// LoginService 定义了登录服务的结构体。
// 包含用于服务的上下文，可用于携带请求范围的值、截止时间等。
type LoginService struct {
	ctx context.Context
}

// NewLoginService 创建一个新的 LoginService 实例。
// 参数：
//
//	ctx: 服务的上下文，通常用于携带请求范围的值或截止时间。
//
// 返回：
//
//	*LoginService: 登录服务的实例。
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run 处理用户的登录流程。
// 参数：
//
//	req: 包含用户登录所需的信息，如邮箱和密码。
//
// 返回：
//
//	resp: 包含用户ID的登录响应。
//	err: 如果发生错误，则返回相应的错误信息。
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 检查输入参数是否为空
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	// 根据邮箱查询用户信息
	row, err := model.GetByEmail(s.ctx, mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	// 验证密码是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// 构建并返回登录成功的响应
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}

	return resp, nil
}
