package model

import (
	"context"

	"gorm.io/gorm"
)

// User 表示用户模型，包含用户的基本信息。
type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"` // 用户的电子邮件地址，唯一索引，不能为空。
	PasswordHashed string `gorm:"type:varchar(255) not null"`             // 用户的密码哈希值，不能为空。
}

// TableName 返回用户模型在数据库中的表名。
func (User) TableName() string {
	return "user"
}

// Create 在数据库中创建一个新的用户记录。
// 参数:
//
//	ctx context.Context: 上下文信息。
//	db *gorm.DB: GORM数据库实例。
//	user *User: 要创建的用户对象。
//
// 返回值:
//
//	error: 如果创建过程中发生错误，则返回错误信息；否则返回nil。
func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

// GetByEmail 根据电子邮件地址从数据库中获取用户信息。
// 参数:
//
//	ctx context.Context: 上下文信息。
//	db *gorm.DB: GORM数据库实例。
//	email string: 用户的电子邮件地址。
//
// 返回值:
//
//	*User: 包含用户信息的指针。
//	error: 如果查询过程中发生错误，则返回错误信息；否则返回nil。
func GetByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
