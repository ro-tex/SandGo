package main

import "github.com/gin-gonic/gin"

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "index",
	})
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", root)

	r.GET("/hello", hello)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
