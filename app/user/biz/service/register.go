package service

import (
	"context"
	"errors"
	"github.com/trashwbin/gomall-demo/app/user/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/user/biz/model"
	user "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

// RegisterService 定义了注册服务的结构体。
// 包含用于服务的上下文，可用于携带请求范围的值、截止时间等。
type RegisterService struct {
	ctx context.Context
}

// NewRegisterService 创建一个新的 RegisterService 实例。
// 参数：
//
//	ctx: 服务的上下文，通常用于携带请求范围的值或截止时间。
//
// 返回：
//
//	*RegisterService: 注册服务的实例。
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run 处理新用户的注册流程。
// 参数：
//
//	req: 包含用户注册所需的信息，如邮箱和密码。
//
// 返回：
//
//	resp: 包含用户ID的注册响应。
//	err: 如果发生错误，则返回相应的错误信息。
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 检查输入参数是否为空
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}

	// 验证密码一致性
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}

	// 对密码进行哈希处理
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建新用户对象
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	// 将新用户保存到数据库
	err = model.Create(s.ctx, mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	// 返回注册成功的响应
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
