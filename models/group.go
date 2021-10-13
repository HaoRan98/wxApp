package models

import "time"

type ProductGroup struct {
	ID			string		`json:"id"`
	ProductID	string		`json:"product_id"`
	Poor		string		`json:"poor"`
	Status		string		`json:"status"`
	CreateTime	time.Time	`json:"create_time"`
	UpdateTime	time.Time	`json:"update_time"`
}

type Member struct {
	ID 				string		`json:"id"`
	ProductGroupID	string 		`json:"product_group_id"`
	UserID			string		`json:"user_id"`
	CreateTime		time.Time	`json:"create_time"`
	UpdateTime		time.Time	`json:"update_time"`
}



