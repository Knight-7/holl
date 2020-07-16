//Author  :     knight
//Date    :     2020/07/14 09:40:16
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     图片上传路由

package routers

import (
	"fmt"
	"holl/dao"
	"holl/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func uploadImage(group *gin.RouterGroup) {
	group.POST("/upload", Authorize(), func(c *gin.Context) {
		var err error
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: err.Error(),
			})
		}
		
		orderID, _ := strconv.ParseInt(c.PostForm("orderId"), 10, 64)
		var count int
		dao.DB.Table("images").Count(&count)
		image := model.Image{
			OrderID: orderID,
			ImageName: strconv.Itoa(count) + ".jpg",
		}
		
		fmt.Println(image, file.Filename)
	})
}

func downloadImage(group *gin.RouterGroup) {
	group.GET("/download", Authorize(), func(c *gin.Context) {

	})
}

//ImageRouter 图片路由注册
func ImageRouter(r *gin.Engine) {
	group := r.Group("/image")
	{
		uploadImage(group)
		downloadImage(group)
	}
}