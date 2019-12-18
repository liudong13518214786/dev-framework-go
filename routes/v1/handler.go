package v1

import (
	"dev-framework-go/conf"
	"dev-framework-go/models"
	s "dev-framework-go/pkg/session"
	"dev-framework-go/pkg/util"
	"fmt"
	"net/http"
	//"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
)

//
//type LoginParm struct {
//	Username string `form:"username"`
//	Password string `form:"password"`
//}

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var res []map[string]interface{}
		r := models.BillRecord(10, 0, "usr_lock2vidladc")
		for index := 0; index < len(r); index++ {
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
				"order_link":            fmt.Sprintf("http://dev-sxs-frontend.mshare.cn/my-orders?order_id=%s", r[index].Uuid),
			}
			res = append(res, tmp)
		}
		//content := fmt.Sprintf("%s%d个", p.OrderDetail.GoodsName, p.OrderDetail.GoodsNum
		c.JSON(200, gin.H{
			"code": 100,
			"data": res,
			"msg":  "ok",
		})
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if username == "" || password == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": conf.MISS_PARAMS,
				"msg":  conf.GetMessage(conf.MISS_PARAMS),
				"data": nil,
			})
			return
		}
		//Useruuid := s.SessionGet(c, "useruuid")
		//if Useruuid == "" {
		s.SessionSet(c, "useruuid", username, conf.COOKIE_EXPIRE_TIME)
		//}
		////todo 一边设置session一边取session,第一次会取不到
		//Useruuid = s.SessionGet(c, "useruuid")
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  username,
			"data": nil,
		})
	}
}

func LogOutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.SessionDel(c)
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  conf.GetMessage(conf.SUCCESS),
			"data": nil,
		})
	}
}
