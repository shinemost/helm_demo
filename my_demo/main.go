package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/query", queryInfo)
	}

	router.Run(":8989")

}

func queryInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"info": "hello world",
		},
	})
}
