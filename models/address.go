package models

import (
	"time"
)

type Address struct {
	ID					string		`json:"id"`
	UserID				string		`json:"user_id"`
	Name				string		`json:"name" gorm:" comment '收件人姓名';" `
	Phone				string		`json:"phone" gorm:" comment '收件人固定电话';"`
	Province			string 		`json:"province" gorm:" comment '省份';"`
	City				string		`json:"city" gorm:" comment '城市';"`
	Area				string		`json:"area" gorm:" comment '县区';"`
	Street				string		`json:"street" gorm:" comment '县区';"`
	Address				string		`json:"address" gorm:" comment '详细地址';"`
	CreateTime			time.Time	`json:"create_time"`
	UpdateTime			time.Time	`json:"update_time"`
}

// 个人地址列表
func GetAddress(userid string) ([]*Address, error) {

	var address = make([]*Address,0)

	if err := DB.Debug().Model(Address{}).Where("user_id = ?", userid).Find(&address).Error; err != nil {
		return address , err
	}

	return address , nil

}

// 单个地址详情
func GetAddressInfo(id string) (*Address , error) {

	var address Address

	if err := DB.Debug().Model(Address{}).Where("id = ?", id).First(&address).Error; err != nil {
		return &address , err
	}

	return &address , nil
}




