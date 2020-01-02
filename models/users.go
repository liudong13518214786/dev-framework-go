package models

import (
	"dev-framework-go/pkg/db"
)

type Users struct {
	Uuid      string
	Email     string
	Real_name string
	Tel       string
	Status    string
}

func GetUserInfoByTel(useruuid string) Users {
	var u Users
	//xx := cache.GetKey("userinfo")
	//if xx != nil {
	//	_ = json.Unmarshal([]byte(xx.(string)), &u)
	//	return u
	//}
	db.DBPool.Table("users").Where("uuid=?", useruuid).First(&u)
	//jsons, _ := json.Marshal(u)
	//cache.SetKey("userinfo", string(jsons), 10)
	return u
}

func GetUserInfoByTelPass(tel, password string) []Users {
	var u []Users
	db.DBPool.Table("users").Where("tel=? and password=?", tel, password).Find(&u)
	return u
}
