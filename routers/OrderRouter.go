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
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//publishorder 发布订单
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

		orderID, err := dao.GetMaxOrderID()
		if err != nil {
			c.JSON(http.StatusOK, Response{
				LoginStatus: true,
				Data: nil,
				ErrorMessage: err.Error(),
			})
			return
		}
		
		publishID := openID.(string)
		deal := model.Deal{
			OrderID: &orderID, 
			PublishID: &publishID,
		}

		orderInfo.OrderID = &orderID
		orderInfo.PublishTime = util.GetLocalTime()

		dao.DB.Create(&deal)
		dao.DB.Create(&orderInfo)

		dao.SetMaxOrderID(orderID)

		c.JSON(http.StatusOK, Response{
			LoginStatus: true,
			Data: orderID,
			ErrorMessage: "",
		})
	})
}

//getPublishOrderByType 根据订单的种类获取发布了单未完成的订单
func getPublishOrderByType(group *gin.RouterGroup) {
	group.GET("/getpublishorderbytype", func(c *gin.Context) {
		log.Println("getpublishorder")
		orderType := c.Query("type")
		var deals []model.Deal
		c.JSON(http.StatusOK, gin.H{
			"orderType": orderType,
			"deals": deals,	
		})
	})
}

//getPublishOrder 根据用户的id获取用户发布的未完成的订单
func getPublishOrder(group *gin.RouterGroup) {
	group.GET("/getpublishorder", Authorize(), func(c *gin.Context) {

	})
}

//getReceiveOrder 根据用户的id获取用户接收的未完成的订单
func getReceiveOrder(group *gin.RouterGroup) {
	group.GET("/getreceiveorder", Authorize(), func(c *gin.Context) {

	})
}

//getHistoryOrder 根据用户的id获取用户的历史订单
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
		getPublishOrderByType(group)
		getPublishOrder(group)
		getReceiveOrder(group)
		getHistoryOrder(group)
		startOrder(group)
		finishOrder(group)
	}
}