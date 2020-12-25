package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

func Server1(ch chan error) {
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

	err := r.Run("x.tar:5000") // 监听并在 0.0.0.0:8080 上启动服务
	ch <- err
}

func Server2(ch chan error) {
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

	err := r.Run("x.tar:5555") // 监听并在 0.0.0.0:8080 上启动服务

	ch <- err
}

func start() {
	ch := make(chan error)

	go Server1(ch)
	go Server2(ch)

	get := <-ch
	log.Println(get)
}

func limitRate() {
	r := rate.NewLimiter(1, 5)
	for {
		if r.AllowN(time.Now(), 2) {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("too many request")
		}
		time.Sleep(time.Second)
	}
}

func MyLimit(next http.Handler) http.Handler {
	r := rate.NewLimiter(1, 5)
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !r.Allow() {
			http.Error(writer, "too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func limitReal() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("OK!!!"))
	})
	_ = http.ListenAndServe(":8888", MyLimit(mux))
}

func main() {
	limitReal()
}
