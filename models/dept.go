package models

import "time"

type SysDept struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	Pid     int             `json:"pid" `
	DeptName string        `json:"deptName" binding:"required"`
	OrderNo int                `json:"orderNo" `
	Level int                `json:"level" `
	Status  int              `json:"status"`
	Remark  string              `json:"remark"`
}
