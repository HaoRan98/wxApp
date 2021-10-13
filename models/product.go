package models

import (
	"fmt"
	"time"
)

// 产品
type Product struct {
	ID				string		`json:"id"`
	Goodstype		string		`json:"goodstype"`
	Name			string		`json:"name" `
	ShortName		string		`json:"short_name" `
	PicUrls			string	`json:"pic_urls" `
	Banner			string	`json:"banner"`
	Sold			int			`json:"sold"`
	Price			int			`json:"price"`
	MinGroupPrice	int			`json:"min_group_price"`
	LinePrice		int			`json:"line_price"`
	CreateTime		time.Time	`json:"create_time"`
	UpdateTime		time.Time	`json:"update_time"`
	Type 			string		`json:"type"`
	SimpleInfo 		string		`json:"simple_info"`
}

// 商品列表
func GetProduct(Goodstype string) ([]*Product, error) {

	var product = make([]*Product,0)

	if err := DB.Debug().Model(Product{}).Where("Goodstype = ?", Goodstype).Find(&product).Error; err != nil {
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

// 单个商品详情
func GetProductInfoFromName(name string) (*Product , error) {

	var product Product

	if err := DB.Debug().Model(Product{}).Where(fmt.Sprintf("name like `%%%s`", name)).First(&product).Error; err != nil {
		return &product , err
	}

	return &product , nil
}

