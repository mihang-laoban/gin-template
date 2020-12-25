package public

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func Server(ch chan error, server string) {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	createParams := []interface{}{
		map[string]interface{}{
			"key":     "1",
			"name":    "John Brown",
			"age":     32,
			"address": "New York No. 1 Lake Park",
			"tags":    []interface{}{"nice", "developer"},
		}, map[string]interface{}{
			"key":     "2",
			"name":    "Jim Green",
			"age":     42,
			"address": "London No. 1 Lake Park",
			"tags":    []interface{}{"loser"},
		}, map[string]interface{}{
			"key":     "3",
			"name":    "Joe Black",
			"age":     32,
			"address": "Sidney No. 1 Lake Park",
			"tags":    []interface{}{"cool", "teacher"},
		},
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, createParams)
	})

	err := r.Run(server) // 监听并在 0.0.0.0:8080 上启动服务

	ch <- err
}

func Run() {
	ch := make(chan error)

	go Server(ch, "x.tar:5000")
	go Server(ch, "x.tar:5555")

	get := <-ch
	log.Println(get)
}
