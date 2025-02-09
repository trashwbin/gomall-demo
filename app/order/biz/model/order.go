package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.Debug().Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}
