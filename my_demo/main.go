package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/query", queryInfo)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/query", queryInfo2)
	}

	router.Run(":8989")

}

func queryInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"info": "hello world v1",
		},
	})
}

func queryInfo2(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"info": "hello world v2",
		},
	})
}
