package v1

import (
	"dev-framework-go/conf"
	"dev-framework-go/middleware/session"
	"dev-framework-go/models"
	"dev-framework-go/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BillExtend struct {
	Resume_order_uuid string `json:"resume_order_uuid,omitempty"`
}

func RecordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("p", "1")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			util.ReturnError(c, conf.INVALID_PARAMS, conf.GetMessage(conf.INVALID_PARAMS), nil)
			return
		}
		offset := (pageInt - 1) * conf.PERNUM
		var res []map[string]interface{}
		r := models.BillRecord(conf.PERNUM, offset, "usr_lock2vidladc")
		if r == nil {
			util.ReturnError(c, conf.SUCCESS, "没有查询到数据", nil)
			return
		}
		for index := 0; index < len(r); index++ {
			var billExtend BillExtend
			if r[index].Extend != "" {
				err = json.Unmarshal([]byte(r[index].Extend), &billExtend)
			}
			order_business_status := ""
			if r[index].Order_status == "pay_success" {
				if conf.BILL_BUSINESS_STATUS[r[index].Business_status] != "" {
					order_business_status = conf.BILL_BUSINESS_STATUS[r[index].Business_status]
				} else {
					order_business_status = conf.BILL_PAY_STATUS["pay_success"]
				}
			} else {
				order_business_status = conf.BILL_PAY_STATUS[r[index].Order_status]
			}
			transaction_time := ""
			if r[index].Order_status == "pay_success" && util.TransTime(r[index].Pay_time) != "" {
				transaction_time = r[index].Pay_time.Format("2006/01/02 15:04")
			}
			if r[index].Order_status == "refunding" {
				transaction_time = r[index].Pay_time.Format("2006/01/02 15:04")
			}
			if r[index].Order_status == "refund" && util.TransTime(r[index].Refund_time) != "" {
				transaction_time = r[index].Refund_time.Format("2006/01/02 15:04")
			}
			content := ""
			for j := 0; j < len(r[index].OrderDetails); j++ {
				content += fmt.Sprintf("%s%d个,", r[index].OrderDetails[j].Goods_name, r[index].OrderDetails[j].Goods_num)
			}
			order_link := ""
			if r[index].Business_type == "resume_optimize" {
				order_link = fmt.Sprintf("http://dev-sxs-frontend.mshare.cn/my-orders?order_id=%s", r[index].Uuid)
			} else {
				order_link = fmt.Sprintf("http://dev-sxs-frontend.mshare.cn/my-orders?order_id=%s", billExtend.Resume_order_uuid)
			}
			tmp := map[string]interface{}{
				"order_no":              r[index].Uuid,
				"order_business_status": order_business_status,
				"business_status":       r[index].Business_status,
				"order_status":          r[index].Order_status,
				"business_type":         conf.GOODS_NAME[r[index].Business_type],
				"total_price":           fmt.Sprintf("%.2f", r[index].Total_price/100),
				"transaction_time":      transaction_time,
				"goods_star":            r[index].GoodsStar,
				"content":               content,
				"order_link":            order_link,
			}
			res = append(res, tmp)
		}
		util.ReturnError(c, conf.SUCCESS, conf.GetMessage(conf.SUCCESS), res)
		return
	}
}

// @Summary 用户登录
// @Description 通过用户的账号和密码登录
// @Accept  json
// @Produce  json
// @Param   username     query    string     true        "用户名"
// @Param   password     query    string     true        "密码"
// @Success 200 {string} string   {"code": 100, "msg": "success", "data":nil}
// @Router /api/v1/login [get]
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		sess := session.Default(c)
		if username == "" || password == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": conf.MISS_PARAMS,
				"msg":  conf.GetMessage(conf.MISS_PARAMS),
				"data": nil,
			})
			return
		}
		sess.Options(session.Options{
			Path:     "/",
			Domain:   conf.DOMAIN,
			MaxAge:   conf.COOKIE_EXPIRE_TIME,
			Secure:   false,
			HttpOnly: true,
		})
		sess.Set("useruuid", username)
		_ = sess.Save()
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  "登录成功",
			"data": nil,
		})
	}
}

func LogOutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := session.Default(c)
		s.Clear()
		_ = s.Save()
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  conf.GetMessage(conf.SUCCESS),
			"data": nil,
		})
	}
}

func BillDetailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Query("order_no")
		r := models.BillDetail(orderId, "usr_lock2vidladc")

		transaction_time := ""
		if r.Order_status == "pay_success" && util.TransTime(r.Pay_time) != "" {
			transaction_time = r.Pay_time.Format("2006/01/02 15:04")
		}
		if r.Order_status == "refunding" {
			transaction_time = r.Pay_time.Format("2006/01/02 15:04")
		}
		if r.Order_status == "refund" && util.TransTime(r.Refund_time) != "" {
			transaction_time = r.Refund_time.Format("2006/01/02 15:04")
		}
		content := ""
		for j := 0; j < len(r.OrderDetails); j++ {
			content += fmt.Sprintf("%s%d个,", r.OrderDetails[j].Goods_name, r.OrderDetails[j].Goods_num)
		}

		order_business_status := ""
		if r.Order_status == "pay_success" {
			if conf.BILL_BUSINESS_STATUS[r.Business_status] != "" {
				order_business_status = conf.BILL_BUSINESS_STATUS[r.Business_status]
			} else {
				order_business_status = conf.BILL_PAY_STATUS["pay_success"]
			}
		} else {
			order_business_status = conf.BILL_PAY_STATUS[r.Order_status]
		}
		util.ReturnError(c, 100, "ok", gin.H{
			"order_no":              r.Uuid,
			"order_status":          r.Order_status,
			"expire_time":           util.TransTime(r.Expire_time),
			"total_price":           fmt.Sprintf("%.2f", r.Total_price/100),
			"business_status":       r.Business_status,
			"order_business_status": order_business_status,
			"business_type":         conf.GOODS_NAME[r.Business_type],
			"content":               content,
			"goods_star":            r.GoodsStar,
			"transaction_time":      transaction_time,
		})
		return
	}
}

func UserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tel := c.Query("tel")
		res := models.GetUserInfoByTel(tel)
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  "查询成功",
			"data": res,
		})
	}
}
