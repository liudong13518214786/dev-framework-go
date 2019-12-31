package v2

import (
	"dev-framework-go/conf"
	"dev-framework-go/models"
	"dev-framework-go/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func WriteBlogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		img := c.PostForm("img")
		title := c.PostForm("title")
		content := c.PostForm("content")
		if img == "" || title == "" || content == "" {
			util.ReturnError(c, 500, "缺少参数", nil)
			return
		}
		models.WriteBlog(title, img, content, "")
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "success",
			"data": nil,
		})
	}
}
