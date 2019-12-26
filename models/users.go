package models

import "dev-framework-go/pkg/db"

type Users struct {
	Uuid      string
	Email     string
	Real_name string
	Tel       string
	Status    string
}

func GetUserInfoByTel(tel string) Users {
	var u Users
	db.DBPool.Table("users").Select("uuid,email,real_name,tel,status").Where("tel=?", tel).First(&u)
	return u
}
