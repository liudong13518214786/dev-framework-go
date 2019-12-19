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
	Total_price     float64
	GoodsStar       int
	OrderDetails    []OrderDetail
}

type OrderDetail struct {
	Goods_name  string
	Goods_num   int
	Order_id    string
	Goods_price int
	Goods_id    string
}

func BillRecord(limit int, offset int, useruuid string) []Record {
	var r []Record
	d := db.DBPool.Table("bill_record AS o").Select("o.uuid, o.user_uuid,o.order_status,o.business_type," +
		"o.expire_time,o.extend,o.business_status,o.pay_time,o.refund_time,o.total_price,oe.goods_star").Joins("left join " +
		"order_evaluate as oe on o.uuid=oe.order_uuid")
	if useruuid != "" {
		d = d.Where("o.user_uuid=?", useruuid)
	}
	d.Limit(limit).Offset(offset).Order("o.build_time DESC").Find(&r)
	if len(r) == 0 {
		return nil
	}
	var orderIdList []string
	for i := 0; i < len(r); i++ {
		orderIdList = append(orderIdList, r[i].Uuid)
	}
	var od []OrderDetail
	db.DBPool.Table("order_detail").Select("goods_name, goods_num,order_id").Where("order_id in (?)", orderIdList).Find(&od)
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(od); j++ {
			if r[i].Uuid == od[j].Order_id {
				r[i].OrderDetails = append(r[i].OrderDetails, od[j])
			}
		}
	}
	return r
}

func BillDetail(order_no string, useruuid string) Record {
	var r Record
	db.DBPool.Table("bill_record AS br").Select("br.uuid, br.user_uuid, br.business_type, br.order_status,"+
		" br.expire_time, br.extend,br.charge_id, br.business_status,br.total_price,br.pay_time,br.refund_time, "+
		"oe.goods_star").Joins("left join order_evaluate oe on br.uuid=oe.order_uuid").Where("br.uuid=? "+
		"and br.user_uuid=?", order_no, useruuid).Find(&r)
	var d OrderDetail
	var de []OrderDetail
	db.DBPool.Table("order_detail").Select("goods_price, goods_id, goods_num, goods_name, "+
		"build_time").Where("order_id=? and user_uuid=?", order_no, useruuid).Find(&d)
	de = append(de, d)
	r.OrderDetails = de
	return r
}
