// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ShopOrderDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ShopOrderDao struct {
	gmvc.M                   // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB           // DB is the raw underlying database management object.
	Table   string           // Table is the table name of the DAO.
	Columns shopOrderColumns // Columns contains all the columns of Table that for convenient usage.
}

// ShopOrderColumns defines and stores column names for table c_shop_order.
type shopOrderColumns struct {
	Id          string // ID
	OrderSn     string // 订单编号
	UserId      string // 会员ID
	OrderAmount string // 订单金额
	PayAmount   string // 实付金额
	Status      string // 状态 0创建 1已支付 2已发货 3已收货 4已完成 9已取消
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	PaidAt      string // 支付时间
	DueExpireAt string // 应过期时间
	ExpiredAt   string // 过期时间
}

func NewShopOrderDao() *ShopOrderDao {
	return &ShopOrderDao{
		M:     g.DB("default").Model("c_shop_order").Safe(),
		DB:    g.DB("default"),
		Table: "c_shop_order",
		Columns: shopOrderColumns{
			Id:          "id",
			OrderSn:     "order_sn",
			UserId:      "user_id",
			OrderAmount: "order_amount",
			PayAmount:   "pay_amount",
			Status:      "status",
			CreatedAt:   "created_at",
			UpdatedAt:   "updated_at",
			PaidAt:      "paid_at",
			DueExpireAt: "due_expire_at",
			ExpiredAt:   "expired_at",
		},
	}
}
