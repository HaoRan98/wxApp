package models

import (
	"time"
)

type Address struct {
	ID					string		`json:"id"`
	UserID				string		`json:"user_id"`
	ReceiverName		string		`json:"receiver_name" gorm:" comment '收件人姓名';" `
	ReceiverPhone		string		`json:"receiver_phone" gorm:" comment '收件人固定电话';"`
	ReceiverMobile		string		`json:"receiver_mobile" gorm:" comment '收件人移动电话';"`
	ReceiverProvince	string 		`json:"receiver_province" gorm:" comment '省份';"`
	ReceiverCity		string		`json:"receiver_city" gorm:" comment '城市';"`
	ReceiverDistrict	string		`json:"receiver_district" gorm:" comment '县区';"`
	ReceicerAddress		string		`json:"receicer_address" gorm:" comment '详细地址';"`
	ReceiverZip			string		`json:"receiver_zip" gorm:" comment '邮编';"`
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




