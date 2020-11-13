package routers

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(group *gin.RouterGroup) {
	group.GET("/h", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
}

func HelloRouter(r *gin.Engine) {
	group := r.Group("/hello")
	{
		hello(group)
	}
}
