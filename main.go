//Author  :     knight
//Date    :     2020/07/11 21:03:22
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :

package main

import (
	"context"
	"holl/config"
	"holl/dao"
	"holl/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func initMysqlAndRedis() {
	var err error
	//获取配置文件
	if err = config.Init(); err != nil {
		log.Println(err)
	}

	//初始化数据库连接
	if err = dao.InitMySQL(); err != nil {
		log.Println(err)
		return
	}
	dao.Migrate()

	//初始化Redis连接池
	dao.InitRedis()

	//初始化oss连接
	if err = dao.InitOss(); err != nil {
		log.Println(err)
		return
	}

	conn := dao.RedisPool.Get()
	r, err := redis.String(conn.Do("get", "user"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(r)
}

func startRouter() {
	//初始化引擎
	r := gin.Default()

	//设置路由
	routers.UserRouter(r)
	routers.OrderRouter(r)
	routers.ImageRouter(r)
	//禁止控制台颜色显示
	gin.DisableConsoleColor()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20

	//启动并静听端口8080
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func main() {
	initMysqlAndRedis()
	defer func() {
		dao.CloseMysql()
		dao.CloseRedis()
	}()

	startRouter()
}