package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"wxApp/pkg/setting"
)

var (
	DB   *gorm.DB
	err  error
)

func Setup() {
	// Initialize database
	DB, err = gorm.Open(setting.DatabaseSetting.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name))



	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	CheckTable()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CheckTable() {
	if !DB.HasTable("sys_user") {
		DB.CreateTable(SysUser{})
	} else {
		DB.AutoMigrate(SysUser{})
	}
	if !DB.HasTable("sys_dept") {
		DB.CreateTable(SysDept{})
	} else {
		DB.AutoMigrate(SysDept{})
	}
	if !DB.HasTable("sys_role") {
		DB.CreateTable(SysRole{})
	} else {
		DB.AutoMigrate(SysRole{})
	}
	if !DB.HasTable("sys_menu") {
		DB.CreateTable(SysMenu{})
	} else {
		DB.AutoMigrate(SysMenu{})
	}
	if !DB.HasTable("casbin_rule") {
		DB.CreateTable(CasbinRule{})
	} else {
		DB.AutoMigrate(CasbinRule{})
	}
	if !DB.HasTable("address") {
		DB.CreateTable(Address{})
	} else {
		DB.AutoMigrate(Address{})
	}
	if !DB.HasTable("order") {
		DB.CreateTable(Order{})
	} else {
		DB.AutoMigrate(Order{})
	}
	if !DB.HasTable("order_item") {
		DB.CreateTable(OrderItem{})
	} else {
		DB.AutoMigrate(OrderItem{})
	}
	if !DB.HasTable("pay_info") {
		DB.CreateTable(PayInfo{})
	} else {
		DB.AutoMigrate(PayInfo{})
	}
	if !DB.HasTable("product") {
		DB.CreateTable(Product{})
	} else {
		DB.AutoMigrate(Product{})
	}
	if !DB.HasTable("category") {
		DB.CreateTable(Category{})
	} else {
		DB.AutoMigrate(Category{})
	}
	if !DB.HasTable("wx_user") {
		DB.CreateTable(WXUser{})
	} else {
		DB.AutoMigrate(WXUser{})
	}
}

func CheckError(err error) {
	log.Println(err)
}

