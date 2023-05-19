package main

import (
	"github.com/gin-gonic/gin"
	"go-web-test/controller"
	"go-web-test/reposity"
)

func main() {
	if err := reposity.InitData("data/"); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		data := controller.QueryPageInfo(c)
		c.JSON(200, data)
	})
	err := r.Run()
	// 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		return
	}
}
