package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("当前的url->", c.Request.URL.String())
	}
}
func main() {
	engine := gin.New()

	//The middleware effect on all request
	//engine.Use(Test(), gin.Logger())

	engine.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})

	//The middleware effect on this request，which name is hello
	engine.GET("/hello", Test(), func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	//The middleware effect on this route group
	api := engine.Group("/api", Test())
	{
		api.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "api hello")
		})
		api.GET("/world", func(c *gin.Context) {
			c.String(http.StatusOK, "api world")
		})
	}
	engine.Run()
}
