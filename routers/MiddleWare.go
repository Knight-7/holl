//Author  :     knight
//Date    :     2020/07/13 12:45:58
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     中间件的设置

package routers

import (
	"holl/dao"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Response 响应
type Response struct {
	LoginStatus  bool        `json:"login_status"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errMsg"`
}

//Authorize 登陆授权中间件，获取sessionID,然后查找redis缓存中的openID和sessionKey
//如果查找失败，则需要登陆
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("session-id")
		openID, sessionKey, err := dao.GetSessionInfo(sessionID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, Response{
				LoginStatus:  false,
				Data:         nil,
				ErrorMessage: err.Error(),
			})
			c.Abort()
			return
		}

		log.Println(sessionID, openID, sessionKey)
		// 将openID和sessionKey保存到context的key中
		c.Set("open-id", openID)
		c.Set("session-key", sessionKey)

		c.Next()
		return
	}
}

//LoginAuthorize 用于检查登录时的登陆状态，如果已经登陆，那么则直接返回，否则登陆
func LoginAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {	
		sessionID := c.GetHeader("session-id")

		_, _, err := dao.GetSessionInfo(sessionID)
		if err != nil {
			log.Println(err)
			c.Next()
			return
		}

		c.JSON(http.StatusNoContent, LoginResponse{
			StatusCode: 0,
			SessionID: "",
			Error: nil,
		})
		c.Abort()
		return
	}
}
