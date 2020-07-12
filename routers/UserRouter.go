//Author  :     knight
//Date    :     2020/07/12 22:19:32
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     用户的路由

package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLoginRouter(group *gin.RouterGroup) {
	group.GET("/login", func(c *gin.Context) {
		appID := "wxd666797118014c6f"
		appSecret := "2bc0041de520030d4651daea3d2547e6"

		code := c.Query("code")

		requestURL := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
			appID, appSecret, code)

		response, err := http.Get(requestURL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
		defer response.Body.Close()
		
		

		c.JSON(http.StatusOK, "hello")
	})
}

//UserRouter 用户路由
func UserRouter(router *gin.Engine) {
	group := router.Group("/user")
	{
		userLoginRouter(group)
	}
}