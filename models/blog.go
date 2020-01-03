package models

import (
	"dev-framework-go/pkg/db"
	"dev-framework-go/pkg/util"
	"encoding/json"
	"fmt"
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

type TagList struct {
	Tag string
}

type TimeCate struct {
	Num        int
	Build_time string
}

type HotBlog struct {
	Uuid     string
	Title    string
	Read_num int
}

func WriteBlog(title, img_url, info, tag string) string {
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
	return bid
}

func ModifyBlog(bid, title, img_url, info, tag string) {
	var res Blog
	tagList := strings.Split(tag, ";")
	tagStr, _ := json.Marshal(tagList)
	db1 := db.DBPool.Table("blogs").Where("uuid=?", bid).First(&res)
	db1.Model(&res).Updates(map[string]interface{}{"title": title, "img_url": img_url, "info": info, "tag": string(tagStr)})
}

func GetBlogByid(bid string) []Blog {
	var res []Blog
	db.DBPool.Table("blogs").Where("uuid=?", bid).Find(&res)
	return res
}

func GetBlog(limit, offset int, keyword, stype string) []Blog {
	var res []Blog
	db1 := db.DBPool.Table("blogs").Order("build_time DESC").Limit(limit).Offset(offset)
	if keyword != "" {
		if stype == "info" {
			s := "%" + keyword + "%"
			db1 = db1.Where("info LIKE ?", s)
		} else if stype == "tag" {
			y := fmt.Sprintf("tag ? '%s'", keyword)
			db1 = db1.Where(y)
		} else {
			db1 = db1.Where("to_char(build_time, 'YYYY-MM-DD')=?", keyword)
		}

	}
	db1.Find(&res)
	return res
}

func GetTotalNum(keyword string) int {
	var count int
	db1 := db.DBPool.Table("blogs")
	if keyword != "" {
		db1 = db1.Where("info like ?", "%"+keyword+"%")
	}
	db1.Count(&count)
	return count
}

func DetailBlog(uuid string) Blog {
	var res Blog
	db.DBPool.Table("blogs").Where("uuid=?", uuid).Find(&res)
	return res
}

func UpdateNum(uuid string) {
	var res Blog
	db1 := db.DBPool.Table("blogs").Where("uuid=?", uuid).First(&res)
	db1.Model(&res).Update("read_num", res.ReadNum+1)
}

func GetBlogTag() []string {
	var res TagList
	db.DBPool.Table("blogs").Select("json_agg(tag) as Tag").Find(&res)
	var tagArr [][]string
	_ = json.Unmarshal([]byte(res.Tag), &tagArr)
	var result []string
	for i := 0; i < len(tagArr); i++ {
		for j := 0; j < len(tagArr[i]); j++ {
			result = append(result, tagArr[i][j])
		}
	}
	result = util.RepeatArr(result)
	return result[:12]
}

func GetCateByTime() []TimeCate {
	var res []TimeCate
	db.DBPool.Table("blogs").Select("count(1) as num, to_char(build_time, 'YYYY-MM-DD') as build_time").
		Group("to_char(build_time, 'YYYY-MM-DD')").Order("num desc").Limit(12).Find(&res)
	return res
}

func GetBlogByReadNum() []HotBlog {
	var res []HotBlog
	db.DBPool.Table("blogs").Select("uuid,title, read_num").Order("read_num DESC").Limit(5).Find(&res)
	return res
}
