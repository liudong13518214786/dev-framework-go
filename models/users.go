package models

import (
	"dev-framework-go/pkg/db"
	"fmt"
)

var Users struct {
	uuid string
	//paytime time.Time
}

func BillRecord() {
	db.DBPool.Limit(1).Find(&Users)
	fmt.Println(Users.uuid)
}
