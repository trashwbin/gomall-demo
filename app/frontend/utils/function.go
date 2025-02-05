package utils

import "context"

// GetUserIdFromCtx 从上下文 context 中获取用户ID。
// 该函数主要用于从给定的上下文（context.Context）中提取用户ID信息。
// 参数:
//
//	ctx context.Context: 包含用户ID的上下文。
//
// 返回值:
//
//	int32: 用户ID。如果没有找到用户ID，则返回0。
func GetUserIdFromCtx(ctx context.Context) int32 {
	// 从上下文中获取与SessionUserId键关联的值。
	userId := ctx.Value(SessionUserId)

	// 如果没有找到用户ID，或者用户ID为nil，则返回默认值0。
	if userId == nil {
		return 0
	}

	// 将获取到的用户ID断言为int32类型，并返回。
	return userId.(int32)
}
