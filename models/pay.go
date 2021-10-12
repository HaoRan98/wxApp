package models

import "time"

type PayInfo struct {
	UserID			string		`json:"user_id"`
	OrderNo			string		`json:"order_no"`
	PayPlatform		string		`json:"pay_platform" gorm:" comment '支付平台：1-支付宝，2-微信';"`
	PlatformNumber	string		`json:"platform_number" gorm:" comment '支付流水号';"`
	PlatformStatus	string		`json:"platform_status" gorm:" comment '支付状态';"`
	CreateTime		time.Time 	`json:"create_time"`
	UpdateTime		time.Time	`json:"update_time"`
}


