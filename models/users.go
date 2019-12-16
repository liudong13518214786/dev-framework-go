package models

import (
	"database/sql"
	"dev-framework-go/pkg/db"
	"fmt"
	"github.com/wonderivan/logger"
)

//测试数据库
type Project struct {
	Uuid            sql.NullString
	User_uuid       sql.NullString
	Order_status    sql.NullString
	Business_type   sql.NullString
	Build_time      sql.NullString
	Expire_time     sql.NullString
	Extend          sql.NullString
	Business_status sql.NullString
	Pay_time        sql.NullString
	Refund_time     sql.NullString
	Total_price     sql.NullString
	Goods_star      sql.NullString
}

//查询单条
func GetBillList(limit, offset, user_uuid string) []*Project {
	var res []*Project
	sqls := "SELECT o.uuid, o.user_uuid,o.order_status,o.business_type, o.build_time,o.expire_time," +
		"o.extend,o.business_status,o.pay_time,o.refund_time,o.total_price," +
		"oe.goods_star FROM bill_record AS o LEFT JOIN order_evaluate AS oe ON(o.uuid=oe.order_uuid) "
	where := " WHERE (o.order_status!='not_pay' or o.expire_time>'now()') AND (o.business_type!='resume_optimize_rework' or o.order_status!='not_pay' )"
	if user_uuid != "" {
		where += fmt.Sprintf(" AND o.user_uuid='%s'", user_uuid)
	}
	order := " ORDER BY o.build_time DESC"
	cond := fmt.Sprintf(" limit %s offset %s", limit, offset)
	rows, err := db.DBPool.Query(sqls + where + order + cond)
	defer rows.Close()
	if err != nil {
		logger.Error(err)
		return res
	}
	for rows.Next() {
		var p Project
		err := rows.Scan(&p.Uuid, &p.User_uuid, &p.Order_status, &p.Business_type, &p.Build_time, &p.Expire_time,
			&p.Extend, &p.Business_status, &p.Pay_time, &p.Refund_time, &p.Total_price, &p.Goods_star)

		if err != nil {
			fmt.Println(err)
		}
		res = append(res, &p)
	}
	return res
}
