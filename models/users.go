package models

import (
	"dev-framework-go/pkg/db"
	"time"
)

type Record struct {
	Uuid            string //必须大写
	User_uuid       string
	Order_status    string
	Business_type   string
	Expire_time     time.Time
	Extend          string
	Business_status string
	Pay_time        time.Time
	Refund_time     time.Time
	Total_price     int64
	GoodsStar       int64
}

//type BillExtend struct {
//	Resume_order_uuid string
//}

func BillRecord(limit int, offset int, useruuid string) []Record {
	var r []Record
	//db.DBPool.Table("bill_record").Where("uuid = ?", "usr_bxw8izhbbcgj").First(&User)
	d := db.DBPool.Table("bill_record AS o").Select("o.uuid, o.user_uuid,o.order_status,o.business_type," +
		"o.expire_time,o.extend,o.business_status,o.pay_time,o.refund_time,o.total_price,oe.goods_star").Joins("left join " +
		"order_evaluate as oe on o.uuid=oe.order_uuid")
	if useruuid != "" {
		d = d.Where("o.user_uuid=?", useruuid)
	}
	d.Limit(limit).Offset(offset).Find(&r)
	//db.DBPool.Table("order_detail").Select("goods_name, goods_num")
	return r
}
