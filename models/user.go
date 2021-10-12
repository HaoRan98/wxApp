package models

import (
	"log"
	"time"
)

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	UserAccount string `json:"user_account"`
	DepID       string `json:"dep_id"`
	DepartName  string `json:"depart_name"`
	ParentID    string `json:"parent_id"`
}

type SysUser struct {
	ID        	string      `gorm:"primary_key" json:"id"`
	Username  	string   	`json:"username" gorm:"unique_index;not null"`
	Nickname  	string    	`json:"nickname"`
	Password  	string    	`json:"password"`
	AvatarUrl 	string    	`json:"avatar_url" gorm:"default:'static/upload/avatar/default.png'"`
	AddressID	string 		`json:"address_id"`
	RoleId    	int       	`json:"role_id" `
	Status      int     	`json:"status" gorm:"default:1"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	Dept      	int       	`json:"dept_id" `
	Phone      	string      `json:"phone" `
	Email      	string      `json:"email" `
	Remark  string  		`json:"remark" `
	SysRole SysRole 		`gorm:"ForeignKey:RoleId" json:"role"`
}

type WXUser struct {
	ID			string		`json:"id"`
	UserID		string		`json:"user_id"`
	OpenID		string		`json:"open_id"`
	Union		string		`json:"union"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func CreateData(data interface{}) error {
	if err := DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func EditData(data interface{}) error {
	err := DB.Debug().Updates(&data).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteData(data interface{}) error {
	err := DB.Debug().Delete(&data).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUserInfo(id string) (*SysUser , error) {

	var user SysUser

	if err := DB.Debug().Model(SysUser{}).Where("id = ?", id).First(&user).Error;err != nil {
		return nil, err
	}

	return &user , nil

}

