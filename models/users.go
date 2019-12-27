package models

import (
	"dev-framework-go/pkg/cache"
	"dev-framework-go/pkg/db"
	"encoding/json"
)

type Users struct {
	Uuid      string
	Email     string
	Real_name string
	Tel       string
	Status    string
}

func GetUserInfoByTel(tel string) Users {
	var u Users
	xx := cache.GetKey("userinfo")
	if xx != nil {
		_ = json.Unmarshal([]byte(xx.(string)), &u)
		return u
	}
	db.DBPool.Table("users").Select("uuid,email,real_name,tel,status").Where("tel=?", tel).First(&u)
	jsons, _ := json.Marshal(u)
	cache.SetKey("userinfo", string(jsons), 10)
	return u
}
