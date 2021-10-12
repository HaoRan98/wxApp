package models

import "time"

type SysMenu struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	Name string             `json:"name" `
	Path string             `json:"path"`
	Sort int                `json:"sort"`
	MenuMeta
	Component string    `json:"component" gorm:"column:component"`
	Type   int             `json:"type"`
	Level   int             `json:"level"`
	ParentId  int           `json:"parent_id"`
	Status  int              `json:"status"`
	ApiPath string         `json:"apiPath"`
	ApiMethod int       `json:"apiMethod"`
	Keepalive int            `json:"keepalive"`
	IsExt    int              `json:"is_ext"`
}
type MenuMeta struct {
	Flag int              `json:"flag"  gorm:"column:flag"`
	Affix  bool             `json:"affix"`
	Icon string             `json:"icon"`
	EnTitle      string     `json:"en_title"`
	CnTitle    string         `json:"cn_title" gorm:"column:cn_title"`
}
