package models

import (
	"dev-framework-go/pkg/db"
	"fmt"
)

//测试数据库
type Project struct {
}

//查询单条
func GetBillList(limit, offset, user_uuid string) *Project {
	var p Project
	sqls := "SELECT o.uuid, o.user_uuid,o.order_status,o.business_type, o.build_time,o.expire_time," +
		"o.cancle_time,o.extend,o.business_status,o.pay_time,o.refund_time,o.total_price, o.user_destroy," +
		"o.update_time, oe.goods_star FROM bill_record AS o LEFT JOIN order_evaluate AS oe ON(o.uuid=oe.order_uuid) "
	where := " WHERE (o.order_status!='not_pay' or o.expire_time>'now()') AND (o.business_type!='resume_optimize_rework' or o.order_status!='not_pay' )"
	if user_uuid != "" {
		where += fmt.Sprintf(" AND o.user_uuid=%s", user_uuid)
	}
	order := " ORDER BY o.build_time DESC"
	cond := fmt.Sprintf(" limit %s offset %s", limit, offset)
	rows := db.DBPool.QueryRow(sqls + where + order + cond)
	err := rows.Scan()
	if err != nil {

	}
	return &p
}
