package conf

var BILL_PAY_STATUS = map[string]string{
	"not_pay":     "未支付",
	"pay_success": "支付成功",
	"refund":      "退款完成",
	"refunding":   "退款中",
}

var BILL_BUSINESS_STATUS = map[string]string{
	"committed": "已提交简历",
	"doing":     "订单处理中",
	"finish":    "订单完成",
	"replace":   "简历已替换",
}

var GOODS_NAME = map[string]string{
	"resume_optimize":        "简历优化",
	"resume_optimize_rework": "简历优化返工次数",
}
