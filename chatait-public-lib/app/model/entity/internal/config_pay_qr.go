// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// ConfigPayQr is the golang structure for table c_config_pay_qr.
type ConfigPayQr struct {
	Id     int    `orm:"id,primary"    json:"id"`     // ID
	Amount int    `orm:"amount,unique" json:"amount"` // 金额
	PayUrl string `orm:"pay_url"       json:"payUrl"` // 二维码
}
