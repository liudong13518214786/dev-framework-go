package models

import (
	"dev-framework-go/pkg/db"
	"dev-framework-go/pkg/util"
)

type Blog struct {
	Uuid       string
	Useruuid   string
	Title      string
	Img_url    string
	Info       string
	Tag        string
	Build_time string
	ReadNum    int
}

func WriteBlog(title, img_url, info, tag string) {
	nowTime := util.GetNowTime()
	bid := util.GenerateRandomString("bid", 8)
	b := Blog{
		Uuid:       bid,
		Useruuid:   "",
		Title:      title,
		Img_url:    img_url,
		Info:       info,
		Tag:        tag,
		Build_time: nowTime,
	}
	db.DBPool.Create(&b)
}
