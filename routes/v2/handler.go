package v2

import (
	"dev-framework-go/conf"
	"dev-framework-go/models"
	"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		err := c.SaveUploadedFile(file, conf.UploadDst+file.Filename)
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
		models.WriteBlog(title, img, content, tags)
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "success",
			"data": nil,
		})
	}
}

func GetBlogListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("p", "1")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			util.ReturnError(c, conf.INVALID_PARAMS, conf.GetMessage(conf.INVALID_PARAMS), nil)
			return
		}
		offset := (pageInt - 1) * conf.PERNUM
		res := models.GetBlog(conf.PERNUM, offset)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "success",
			"data": res,
		})
	}
}
