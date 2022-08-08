package models

import "time"

// 订单表
type Order struct {
	ID			string		`json:"id"`
	OrderSN		string 		`json:"order_sn"`
	UserID		string		`json:"user_id" `
	TotalPrice	int 		`json:"total_price"`
	ExpressType	int 		`json:"express_type"`
	numbers 	int 		`json:"numbers"`
	Postage		int			`json:"postage" gorm:" comment '运费';"`
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
	OrderSN				int64		`json:"order_sn" gorm:" comment '订单号';"`
	ProductID			string		`json:"product_id"`
	ProductName			string		`json:"product_name"`
	ProductImage 		string		`json:"product_image"`
	CurrentUnitPrice	int			`json:"current_unit_price" gorm:" comment '生成订单的商品单价';"`
	Quantity			int			`json:"quantity" gorm:" comment '商品数量';"`
	TotalPrice  		int			`json:"total_price" gorm:" comment '商品总价';"`
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



