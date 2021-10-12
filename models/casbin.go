package models

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"strings"
)

type CasbinRule struct {
	Ptype       string `json:"ptype" gorm:"column:p_type"`
	RoleId      string `json:"rolen_id" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
}

var (
	Casbins Casbin
	CASBIN  *casbin.Enforcer
)

func InitCasbin()  {
	a := gormadapter.NewAdapterByDB(DB)
	e := casbin.NewEnforcer(Casbins.ModelPath, a)
	e.EnableLog(true)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)

	CASBIN =e
	return
}


func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}


func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}
