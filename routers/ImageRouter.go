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
	"log"
	"net/http"
	"os"
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
		imageNum, err := dao.GetMaxImageNum()
		if err != nil {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: err.Error(),
			})
		}

		imageName := strconv.FormatInt(imageNum, 10) + ".jpg"
		c.SaveUploadedFile(file, "tmp/" + file.Filename)
		if err = dao.OssUploadFile("holl/" + imageName, "tmp/" + file.Filename); err != nil {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: err.Error(),
			})
		}

		image := model.Image{
			OrderID: &orderID,
			ImageName: &imageName,
		}

		dao.DB.Create(&image)
		dao.SetMaxImageNum(imageNum)

		if err = os.Remove("tmp/" + file.Filename); err != nil {
			log.Println("Delete tmprary file failed, filename:", imageName)
		}
		
		fmt.Println(image, file.Filename)
	})
}

//ImageRouter 图片路由注册
func ImageRouter(r *gin.Engine) {
	group := r.Group("/image")
	{
		uploadImage(group)
	}
}