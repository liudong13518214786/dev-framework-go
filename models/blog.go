package models

import (
	"dev-framework-go/pkg/db"
	"dev-framework-go/pkg/util"
	"encoding/json"
	"strings"
	"time"
)

type Blog struct {
	Uuid       string
	Useruuid   interface{}
	Title      string
	Img_url    string
	Info       string
	Tag        string
	Build_time time.Time
	ReadNum    int
}

func WriteBlog(title, img_url, info, tag string) {
	//nowTime:=util.GetNowTime()
	bid := util.GenerateRandomString("bid", 8)
	tagList := strings.Split(tag, ";")
	tagStr, _ := json.Marshal(tagList)
	b := Blog{
		Uuid:       bid,
		Useruuid:   nil,
		Title:      title,
		Img_url:    img_url,
		Info:       info,
		Tag:        string(tagStr),
		Build_time: time.Now(),
	}
	db.DBPool.Create(&b)
}

func GetBlog(limit, offset int) []Blog {
	var res []Blog
	db.DBPool.Table("blogs").Limit(limit).Offset(offset).Find(&res)
	return res
}
