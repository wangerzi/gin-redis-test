package main

import (
	"fmt"
	"gin-redis-test/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/redis/:addr/:password/:db", func(c *gin.Context) {
		status := true
		message := "success"
		data := gin.H{}

		var params redis.UriCheckParams

		if err := c.ShouldBindUri(&params); err != nil {
			status = false
			message = err.Error()
		} else {
			information, err := redis.GetRdbInfo(params)
			if err != nil {
				status = false
				message = err.Error()
			} else {
				data["info"] = information
			}
		}

		c.JSON(200, gin.H{
			"status":  status,
			"data":    data,
			"message": message,
		})
	})

	err := r.Run(":80")
	if err != nil {
		fmt.Println("serve err:", err)
	}
}
