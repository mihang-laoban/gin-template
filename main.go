package main

import (
	"fmt"
	"gin-template/lower"
	"github.com/gin-gonic/gin"
	"net/http"
)

func start() {
	//LimitReal()
	//Run()
	lower.TestLowerAsync()
	//lower.TestLowerSync()
}

func runGin() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": "ok",
			"nick":    "good",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		type params struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
		var param params
		if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{
			"status": "posted",
			"ids":    param.Id,
			"names":  param.Name,
		})
	})
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

//todo consul register services

func main() {
	runGin()
}
