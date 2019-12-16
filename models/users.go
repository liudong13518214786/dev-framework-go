package models

import (
	"database/sql"
	"dev-framework-go/conf"
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
	Total_price     sql.NullFloat64
	Goods_star      sql.NullInt64
	OrderDetail     OrderDetail
}

type OrderDetail struct {
	GoodsNum  int64
	GoodsName string
}

//查询单条
func GetBillList(limit, offset, user_uuid string) []map[string]interface{} {
	var res []map[string]interface{}
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
		if err == sql.ErrNoRows {
			logger.Debug("没有订单列表数据")
			return res
		}
		logger.Error(err)
		return res
	}
	for rows.Next() {
		var p Project
		var d OrderDetail
		err := rows.Scan(&p.Uuid, &p.User_uuid, &p.Order_status, &p.Business_type, &p.Build_time, &p.Expire_time,
			&p.Extend, &p.Business_status, &p.Pay_time, &p.Refund_time, &p.Total_price, &p.Goods_star)
		if err != nil {
			logger.Error("查询订单列表错误", err)
		}
		detailSql := db.SelectSql("order_detail", []string{"goods_name", "goods_num"}, map[string]interface{}{"order_id": p.Uuid.String}, "order by build_time DESC", "", "")
		detailRow := db.DBPool.QueryRow(detailSql)
		err = detailRow.Scan(&d.GoodsName, &d.GoodsNum)
		if err != nil {
			logger.Error("查询订单详情错误", err)
		}
		p.OrderDetail = d

		order_business_status := ""
		if p.Order_status.String == "pay_success" {
			if conf.BILL_BUSINESS_STATUS[p.Business_status.String] != "" {
				order_business_status = conf.BILL_BUSINESS_STATUS[p.Business_status.String]
			} else {
				order_business_status = conf.BILL_PAY_STATUS["pay_success"]
			}
		} else {
			order_business_status = conf.BILL_PAY_STATUS[p.Order_status.String]
		}
		transaction_time := ""
		if p.Order_status.String == "pay_success" && p.Pay_time.String != "" {
			transaction_time = ""
		}
		if p.Order_status.String == "refunding" {
			transaction_time = ""
		}
		if p.Order_status.String == "refund" && p.Refund_time.String != "" {
			transaction_time = ""
		}
		content := fmt.Sprintf("%s%d个", p.OrderDetail.GoodsName, p.OrderDetail.GoodsNum)

		tmp := map[string]interface{}{
			"order_no":              p.Uuid.String,
			"order_business_status": order_business_status,
			"business_status":       p.Business_status.String,
			"order_status":          p.Order_status.String,
			"business_type":         conf.GOODS_NAME[p.Business_type.String],
			"total_price":           fmt.Sprintf("%.2f", p.Total_price.Float64/100),
			"transaction_time":      transaction_time,
			"goods_star":            p.Goods_star.Int64,
			"content":               content,
			"order_link":            fmt.Sprintf("http://dev-sxs-frontend.mshare.cn/my-orders?order_id=%s", p.Uuid.String),
		}
		res = append(res, tmp)
	}
	return res
}
