package v1

import (
	"dev-framework-go/conf"
	"dev-framework-go/models"
	s "dev-framework-go/pkg/session"
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
		var result []map[string]interface{}
		//Useruuid := s.SessionGet(c, "useruuid")
		res := models.GetBillList("10", "0", "")
		for index := 0; index < len(res); index++ {
			order_business_status := ""
			if res[index].Order_status.String == "pay_success" {
				if conf.BILL_BUSINESS_STATUS[res[index].Business_status.String] != "" {
					order_business_status = conf.BILL_BUSINESS_STATUS[res[index].Business_status.String]
				} else {
					order_business_status = conf.BILL_PAY_STATUS["pay_success"]
				}
			} else {
				order_business_status = conf.BILL_PAY_STATUS[res[index].Order_status.String]
			}
			transaction_time := ""
			if res[index].Order_status.String == "pay_success" && res[index].Pay_time.String != "" {
				transaction_time = ""
			}
			if res[index].Order_status.String == "refunding" {
				transaction_time = ""
			}
			if res[index].Order_status.String == "refund" && res[index].Refund_time.String != "" {
				transaction_time = ""
			}
			tmp := map[string]interface{}{
				"order_no":              res[index].Uuid.String,
				"order_business_status": order_business_status,
				"business_status":       res[index].Business_status.String,
				"order_status":          res[index].Order_status.String,
				"business_type":         conf.GOODS_NAME[res[index].Business_type.String],
				"total_price":           fmt.Sprintf("%.2f", res[index].Total_price.Float64/100),
				"transaction_time":      transaction_time,
				"goods_star":            res[index].Goods_star.String,
				"content":               "",
				"order_link":            "",
			}
			result = append(result, tmp)
		}
		c.JSON(200, gin.H{
			"code": 100,
			"data": result,
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
