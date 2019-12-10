package util

import (
	"crypto/md5"
	"dev-framework-go/conf"
	"encoding/hex"
	"fmt"
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
	if conf.ERRORNOTIFYOPEN == false {
		return nil
	}
	m := gomail.NewMessage()
	//设置发件人
	m.SetHeader("From", conf.SYSTEMEMAILUSER)
	to_user := strings.Split(conf.EMAILTOUSER, ",")
	m.SetHeader("To", to_user...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(conf.SYSTEMEMAILHOST, conf.SYSTEMEMAILPORT, conf.SYSTEMEMAILUSER, conf.SYSTEMEMAILPASS)
	fmt.Println(d)
	err := d.DialAndSend(m)
	fmt.Println(err)
	if err != nil {
		logger.Error("[SEND ERROR EMAIL FAIL]", err)
	}
	return err
}

func GetNowTime() string {
	now := time.Now()
	return now.Format(conf.TIME_FORMAT)
}
