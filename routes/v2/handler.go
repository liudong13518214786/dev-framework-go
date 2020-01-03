package v2

import (
	"dev-framework-go/conf"
	"dev-framework-go/models"
	"dev-framework-go/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 上传
// @Description 上传博客头图
// @Accept  json
// @Produce  json
// @Param   username     query    string     true        "用户名"
// @Success 200 {string} string   {"code": 100, "msg": "上传成功", "data":"图片地址"}
// @Router /api/v1/upload [post]
func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		err := c.SaveUploadedFile(file, conf.UploadDst+file.Filename)
		fmt.Println(err)
		if err != nil {
			util.ReturnError(c, 200, "上传失败", nil)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"data": conf.PicUrlHost + file.Filename,
			"msg":  "上传成功",
		})
		return
	}
}

// @Summary 编写博客
// @Description 如果博客存在，那么则修改
// @Accept  json
// @Produce  json
// @Param   img     query    string     true        "头图"
// @Param   title     query    string     true        "标题"
// @Param   content     query    string     true        "内容"
// @Param   tags     query    string     true        "标签"
// @Param   bid     query    string     false        "博客id"
// @Success 200 {string} string   {"code": 100, "msg": "success", "data":nil}
// @Router /api/v1/write [post]
func WriteBlogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		img := c.PostForm("img")
		title := c.PostForm("title")
		content := c.PostForm("content")
		tags := c.PostForm("tags")
		if img == "" || title == "" || content == "" || tags == "" {
			util.ReturnError(c, 500, "缺少参数", nil)
			return
		}
		bid := c.DefaultPostForm("bid", "")
		if bid == "" {
			bid = models.WriteBlog(title, img, content, tags)
		} else {
			hasB := models.GetBlogByid(bid)
			if len(hasB) == 0 {
				util.ReturnError(c, 500, "参数错误", nil)
				return
			}
			models.ModifyBlog(bid, title, img, content, tags)
		}

		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "success",
			"data": bid,
		})
	}
}

// @Summary 首页
// @Description 获取博客列表
// @Accept  json
// @Produce  json
// @Param   p     query    string     false        "页数"
// @Param   kw     query    string     false        "关键词"
// @Param   stype     query    string     false        "搜索类型"
// @Success 200 {string} string   {"code": 100, "msg": "success", "data":[]}
// @Router /api/v1/blog [get]
func GetBlogListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("p", "1")
		keyword := c.DefaultQuery("kw", "")
		stype := c.DefaultQuery("stype", "info")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			util.ReturnError(c, conf.INVALID_PARAMS, conf.GetMessage(conf.INVALID_PARAMS), nil)
			return
		}
		offset := (pageInt - 1) * conf.PERNUM
		res := models.GetBlog(conf.PERNUM, offset, keyword, stype)
		countNum := models.GetTotalNum(keyword)
		var result []map[string]interface{}
		for i := 0; i < len(res); i++ {
			var tagList []string
			_ = json.Unmarshal([]byte(res[i].Tag), &tagList)
			tmp := map[string]interface{}{
				"title":      res[i].Title,
				"uuid":       res[i].Uuid,
				"img":        res[i].Img_url,
				"info":       res[i].Info,
				"tag":        tagList,
				"build_time": util.TransTime(res[i].Build_time),
				"readnum":    res[i].ReadNum,
			}
			result = append(result, tmp)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "success",
			"data": result,
			"page": map[string]interface{}{
				"page_size":   conf.PERNUM,
				"total":       countNum,
				"currentpage": pageInt,
			},
		})
	}
}

// @Summary 详情
// @Description 博客详情
// @Accept  json
// @Produce  json
// @Param   uuid     query    string     true        "博客id"
// @Success 200 {string} string   {"code": 100, "msg": "success", "data":nil}
// @Router /api/v1/detail [get]
func DetailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Query("uuid")
		if uuid == "" {
			util.ReturnError(c, 500, "缺少参数", nil)
			return
		}
		res := models.DetailBlog(uuid)
		if &res == nil {
			util.ReturnError(c, 500, "参数错误", nil)
			return
		}
		var tagList []string
		_ = json.Unmarshal([]byte(res.Tag), &tagList)
		result := map[string]interface{}{
			"uuid":       res.Uuid,
			"title":      res.Title,
			"build_time": util.TransTime(res.Build_time),
			"read_num":   res.ReadNum,
			"info":       res.Info,
			"tag":        tagList,
			"img_url":    res.Img_url,
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "success",
			"data": result,
		})
		models.UpdateNum(uuid)
	}
}

// @Summary 分类
// @Description 博客的分类
// @Accept  json
// @Produce  json
// @Success 200 {string} string   {"code": 100, "msg": "success", "data":[]}
// @Router /api/v1/cate [get]
func GetClassHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tag := models.GetBlogTag()
		cate := models.GetCateByTime()
		hot := models.GetBlogByReadNum()
		res := map[string]interface{}{
			"tag":  tag,
			"cate": cate,
			"hot":  hot,
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "success",
			"data": res,
		})
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}
