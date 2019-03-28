
package main

import "github.com/gin-gonic/gin"
import "net/http"

// func1: 处理最基本的GET
func func1 (c *gin.Context)  {
    // 回复一个200OK,在client的http-get的resp的body中获取数据
    c.String(http.StatusOK, "test1 OK")
}
// func2: 处理最基本的POST
func func2 (c *gin.Context) {
    // 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}
/*
* C:\Users\App>curl  --data "post by eagle" --request POST http://127.0.0.1:8888/test2
* curl http://127.0.0.1:8888/test1
*
* route group
v1 := router.Group("/v1")
v1.GET("/user/get_username", modules.UserGetUsername)
v2 := router.Group("/v2")
v2.GET("/user/get_username", modules.UserGetUsername)
Is there a nicer solution for that - maybe something like that:

v1_v2 := router.Group("/v1").AnotherGroup("/v2")
v1_v2.GET("/user/get_username", modules.UserGetUsername)

*/

func main(){
    // 注册一个默认的路由器
    router := gin.Default()
    // 最基本的用法
    router.GET("/test1", func1)
	router.POST("/test2", func2)
	
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
	}

	router.GET("/v1/:login", func(c *gin.Context) {
		name := c.Param("login")
		c.String(http.StatusOK, "Hello %s", "V1"+name)
	})

	router.GET("/v2/:login", func(c *gin.Context) {
		name := c.Param("login")
		c.String(http.StatusOK, "Hello %s", "v2"+name)
	})
    // 绑定端口是8888
    router.Run(":8888")
}

func loginEndpoint(c *gin.Context){
	
	user := c.Request.FormValue("user")
	pwd  := c.Request.FormValue("pw")
	c.JSON(http.StatusOK, user + pwd)
}

