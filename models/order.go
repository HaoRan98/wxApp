package models

import "time"

// 订单表
type Order struct {
	ID			string		`json:"id"`
	OrderNo		string		`json:"order_no" gorm:" comment '订单号';"`
	UserID		string		`json:"user_id" `
	Payment		float32 	`json:"payment" gorm:" comment '实际付款金额，单位元，保留2位小数';"`
	PaymentType string		`json:"payment_type" gorm:" comment '支付类型';" `
	Postage		float32		`json:"postage" gorm:" comment '运费';"`
	Status		string		`json:"status" gorm:" comment '订单状态';"`
	PaymentTime time.Time	`json:"payment_time" gorm:" comment '支付时间';"`
	SendTime	time.Time	`json:"send_time" gorm:" comment '发货时间';"`
	EndTime		time.Time	`json:"end_time" gorm:" comment '交易完成时间';"`
	CloseTime   time.Time	`json:"close_time" gorm:" comment '交易关闭时间';"`
	CreateTime	time.Time	`json:"create_time"`
	UpdateTime	time.Time	`json:"update_time"`
}


// 订单明细表
type OrderItem struct {
	ID 					string		`json:"id"`
	UserID				string		`json:"user_id"`
	OrderNo				int64		`json:"order_no" gorm:" comment '订单号';"`
	ProductID			string		`json:"product_id"`
	ProductName			string		`json:"product_name"`
	ProductImage 		string		`json:"product_image"`
	CurrentUnitPrice	float32		`json:"current_unit_price" gorm:" comment '生成订单的商品单价';"`
	Quantity			int			`json:"quantity" gorm:" comment '商品数量';"`
	TotalPrice  		float32		`json:"total_price" gorm:" comment '商品总价';"`
	CreateTime			time.Time	`json:"create_time"`
	UpdateTime			time.Time	`json:"update_time"`
}

func GetOrderInfo(userid string) ([]*Order, error) {

	var order = make([]*Order,0)

	if err := DB.Debug().Model(Order{}).Where("user_id = ?", userid).Find(&order).Error; err != nil {
		return order , err
	}

	return order,nil

}



