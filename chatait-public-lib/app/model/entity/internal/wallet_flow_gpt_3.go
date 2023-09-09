// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// WalletFlowGpt3 is the golang structure for table c_wallet_flow_gpt3.
type WalletFlowGpt3 struct {
	Id         int64  `orm:"id,primary"  json:"id"`         // ID
	UserId     uint64 `orm:"user_id"     json:"userId"`     // 会员ID
	Amount     int    `orm:"amount"      json:"amount"`     // 变动金额
	Total      int    `orm:"total"       json:"total"`      // 变动后的余额
	IsIncr     int    `orm:"is_incr"     json:"isIncr"`     // 增加减少
	TargetType string `orm:"target_type" json:"targetType"` // 目标类型
	TargetId   int64  `orm:"target_id"   json:"targetId"`   // 目标ID
	Remark     string `orm:"remark"      json:"remark"`     // 备注
	AdminName  string `orm:"admin_name"  json:"adminName"`  // 操作管理员
	Year       int    `orm:"year"        json:"year"`       // 年
	Month      int    `orm:"month"       json:"month"`      // 月
	Day        int    `orm:"day"         json:"day"`        // 日
	CreatedAt  int    `orm:"created_at"  json:"createdAt"`  // 创建时间
	UpdatedAt  int    `orm:"updated_at"  json:"updatedAt"`  // 更新时间
}
