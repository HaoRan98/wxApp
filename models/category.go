package models

import "time"

// 分类
type Category struct {
	ID 			string		`json:"id"`
	Name 		string 		`json:"name"`
	ParentID	string		`json:"parent_id"`
	Status		string		`json:"status"`
	SortOrder	int			`json:"sort_order"`
	CreateTime	time.Time	`json:"create_time"`
	UpdateTime	time.Time	`json:"update_time"`
}
