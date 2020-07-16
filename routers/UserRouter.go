//Author  :     knight
//Date    :     2020/07/12 22:19:32
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     用户的路由

package routers

import (
	"encoding/json"
	"fmt"
	"holl/dao"
	"holl/model"
	"holl/util"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserCode 保存用户的OpenID和sessionKey
type UserCode struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

//LoginResponse 登陆返回信息
type LoginResponse struct {
	StatusCode int    `json:"status_code"`
	SessionID  string `json:"session_id"`
	Error      error  `json:"error"`
}

//用户登陆流程
//1、前端登陆调用wx.login()获取code并发送到后端
//2、后端根据code调用微信api接口获得OpenID和SessionKey
//3、后端生成3rd_session与OpenID和SessionKey相对应然后保存到redis缓存中，时效性为2小时
//4、后端在每次登陆时都需要获取用户的信息并更新数据库
//5、后端返回3rd_session给前端，前端将3rd_session保存到缓存中，
//	 每次业务操作时，前段不需要重新登陆，只需要根据将3rd_session发送给后端，
//   后端根据3rd_session找到对应的OpenID和SessionKey
func userLoginRouter(group *gin.RouterGroup) {
	group.POST("/login", LoginAuthorize(),func(c *gin.Context) {
		appID := "wxd666797118014c6f"
		appSecret := "2bc0041de520030d4651daea3d2547e6"

		//获取前段传过来的code
		code := c.PostForm("code")

		//调用微信api获取sessionKey和OpenID
		requestURL := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
			appID, appSecret, code)

		//发送一个请求
		response, err := http.Get(requestURL)
		if err != nil {
			c.JSON(http.StatusBadRequest, LoginResponse{
				StatusCode: http.StatusBadRequest,
				SessionID:  "",
				Error:      err,
			})
			return
		}
		defer response.Body.Close()

		//读取请求的响应
		s, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, LoginResponse{
				StatusCode: http.StatusBadRequest,
				SessionID:  "",
				Error:      err,
			})
			return
		}

		//解析响应字符串到json中
		var userCode UserCode
		json.Unmarshal(s, &userCode)

		//生成sessionID并保存到redis中
		sessionID := util.GenerateSessionID()
		redisConn := dao.RedisPool.Get()
		defer redisConn.Close()

		//在redis缓存中设置session
		if _, err = redisConn.Do("hmset", sessionID, "openID", userCode.OpenID,
			"sessionKey", userCode.SessionKey); err != nil {
			c.JSON(http.StatusBadRequest, LoginResponse{
				StatusCode: http.StatusBadRequest,
				SessionID:  "",
				Error:      err,
			})
			return
		}
		//设置有效时间为2小时
		if _, err = redisConn.Do("expire", sessionID, "7200"); err != nil {
			c.JSON(http.StatusBadRequest, LoginResponse{
				StatusCode: http.StatusBadRequest,
				SessionID:  "",
				Error:      err,
			})
			return
		}

		//向前段发送sessionID
		c.JSON(http.StatusOK, LoginResponse{
			StatusCode: http.StatusOK,
			SessionID:  sessionID,
			Error:      nil,
		})
	})
}

func saveUserInfo(group *gin.RouterGroup) {
	group.POST("/saveuserinfo", Authorize(), func(c *gin.Context) {
		appID := "wxd666797118014c6f"
		var err error

		// 取出open-id
		openID, ok := c.Get("open-id")
		if !ok {
			c.JSON(http.StatusBadRequest, Response{
				LoginStatus:  true,
				Data:         nil,
				ErrorMessage: "get openid failed",
			})
			return
		}
		// 取出session-key
		sessionKey, ok := c.Get("session-key")
		if !ok {
			c.JSON(http.StatusBadRequest, Response{
				LoginStatus:  true,
				Data:         nil,
				ErrorMessage: "get sessionkey failed",
			})
		}
		// 获取表单中的iv偏移量和encryptedData密文
		iv := c.PostForm("iv")
		encryptedData := c.PostForm("encryptedData")
		// 根据sessionKey、iv和encryptedData解析用户数据
		userInfo, err := util.Decrypt(appID, sessionKey.(string), encryptedData, iv)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				LoginStatus:  true,
				Data:         nil,
				ErrorMessage: err.Error(),
			})
		}
		// 将json数据反序列化到user中
		var user model.User
		json.Unmarshal([]byte(userInfo), &user)
		// save：如果数据存在，即更新数据；不存在就添加新的记录
		dao.DB.Save(&user)
		// 向前段返回信息
		c.JSON(http.StatusOK, Response{
			LoginStatus:  true,
			Data:         openID,
			ErrorMessage: "",
		})
	})
}

//UserRouter 用户路由注册
func UserRouter(router *gin.Engine) {
	group := router.Group("/user")
	{
		userLoginRouter(group)
		saveUserInfo(group)
	}
}
