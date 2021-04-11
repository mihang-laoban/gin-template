package main

import (
	"context"
	"fmt"
	"gin-template/lower"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func start() {
	//LimitReal()
	//Run()
	lower.TestLowerAsync()
	//lower.TestLowerSync()
	//todo consul register services
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
			Id   int    `json:"Id"`
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

func Db() {
	type Test struct {
		Id   int
		Name string
	}

	dsn := "root:123@tcp(127.0.0.1:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var test Test
	db.Find(&test)
	//fmt.Println(res.Statement)
	fmt.Println(test.Id, test.Name)
}

func cache() {

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})

	err := rdb.Set(ctx, "a", "tar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	//runGin()
	//Db()
	cache()
}
