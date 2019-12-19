package util

import (
	"crypto/md5"
	"dev-framework-go/conf"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"gopkg.in/gomail.v2"
	"strings"
	"time"
)

func EncodeMD5(signature string) string {
	t := md5.New()
	t.Write([]byte(signature))
	return hex.EncodeToString(t.Sum(nil))
}

func SendMail(subject string, body string) error {
	m := gomail.NewMessage()
	//设置发件人
	m.SetHeader("From", conf.SYSTEMEMAILUSER)
	to_user := strings.Split(conf.EMAILTOUSER, ",")
	m.SetHeader("To", to_user...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(conf.SYSTEMEMAILHOST, conf.SYSTEMEMAILPORT, conf.SYSTEMEMAILUSER, conf.SYSTEMEMAILPASS)
	err := d.DialAndSend(m)
	if err != nil {
		logger.Error("[SEND ERROR EMAIL FAIL]", err)
	}
	return err
}

func GetNowTime() string {
	now := time.Now()
	return now.Format(conf.TIME_FORMAT)
}

func ArrayToSql(r []string) string {
	return fmt.Sprintf("('%s')", strings.Join(r, "','"))
}

func TransTime(t time.Time) string {
	s := t.Format(conf.TIME_FORMAT)
	if s == "0001-01-01 00:00:00" { //golang的默认时间
		return ""
	}
	return s
}
func ReturnError(c *gin.Context, code int, errMsg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  errMsg,
	})
}
