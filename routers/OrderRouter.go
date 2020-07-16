//Author  :     knight
//Date    :     2020/07/13 20:05:20
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     订单的路由

package routers

import (
	"holl/dao"
	"holl/model"
	"holl/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func publishorder(group *gin.RouterGroup) {
	group.POST("/publish", Authorize(), func(c *gin.Context) {
		openID, ok := c.Get("open-id")
		
		if !ok {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: "publishorder: Get openid failed",
			})
			return
		}
		
		var orderInfo model.Order
		if err := c.Bind(&orderInfo); err != nil {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: err.Error(),
			})
			return
		}
		
		orderInfo.Deal = &model.Deal{PublishID: openID.(string)}
		orderInfo.Deal.PublishID = openID.(string)
		orderInfo.PublishTime = util.GetLocalTime()
		dao.DB.Create(&orderInfo)

		c.JSON(http.StatusOK, Response{
			LoginStatus: true,
			Data: orderInfo.ID,
			ErrorMessage: "",
		})
	})
}

func getPublishOrderByType(group *gin.RouterGroup) {
	group.GET("/getpublishorder", Authorize(), func(c *gin.Context) {

	})
}

func getPublishOrder(group *gin.RouterGroup) {
	group.GET("/getpublishorder", Authorize(), func(c *gin.Context) {

	})
}

func getReceiveOrder(group *gin.RouterGroup) {
	group.GET("/getreceiveorder", Authorize(), func(c *gin.Context) {

	})
}

func getHistoryOrder(group *gin.RouterGroup) {
	group.GET("/history", Authorize(), func(c *gin.Context) {
		openID, ok := c.Get("open-id")
		if !ok {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: "",
			})
			return
		}

		var user model.User
		var historyOrders []model.Order
		dao.DB.Table("users").Where("open_id = ?", openID).Model(&user).Related(historyOrders, )
	})
}

func startOrder(group *gin.RouterGroup) {
	group.POST("/start", Authorize(),func(c *gin.Context) {

	})
}

func finishOrder(group *gin.RouterGroup) {
	group.POST("/finish", Authorize(), func(c *gin.Context) {

	})
}

//OrderRouter 订单路由注册
func OrderRouter(router *gin.Engine) {
	group := router.Group("/order")
	{
		publishorder(group)
		getPublishOrder(group)
		getReceiveOrder(group)
		getHistoryOrder(group)
		startOrder(group)
		finishOrder(group)
	}
}