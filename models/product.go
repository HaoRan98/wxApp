package models

import "time"

// 产品
type Product struct {
	ID			string		`json:"id"`
	CategoryID	string		`json:"category_id" gorm:" comment '分类ID';"`
	Name		string		`json:"name" gorm:" comment '名称';"`
	Subtitle	string		`json:"subtitle" gorm:" comment '副标题';"`
	MainImage	string		`json:"main_image" gorm:" comment '主图';"`
	SubImages	string		`json:"sub_images" gorm:" comment '支付流水号';type:varchar(65535);not null"`
	Detail		string		`json:"detail"`
	Price		float32		`json:"price"`
	Stock		int			`json:"stock"`
	Status		string		`json:"status"`
	CreateTime	time.Time	`json:"create_time"`
	UpdateTime	time.Time	`json:"update_time"`
}

// 商品列表
func GetProduct(userid string) ([]*Product, error) {

	var product = make([]*Product,0)

	if err := DB.Debug().Model(Product{}).Where("user_id = ?", userid).Find(&product).Error; err != nil {
		return product , err
	}

	return product , nil

}

// 单个商品详情
func GetProductInfo(id string) (*Product , error) {

	var product Product

	if err := DB.Debug().Model(Product{}).Where("id = ?", id).First(&product).Error; err != nil {
		return &product , err
	}

	return &product , nil
}

