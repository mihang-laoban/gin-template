package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if res.StatusCode != 200 {
	//	fmt.Println("Unexpected status code", res.StatusCode)
	//}
	//
	//// Read the token out of the response body
	//buf := new(bytes.Buffer)
	//io.Copy(buf, res.Body)
	//res.Body.Close()
	//tokenString := strings.TrimSpace(buf.String())
	//
	//// Parse the token
	//token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
	//	// since we only use the one private key to sign the tokens,
	//	// we also only use its public counter part to verify
	//	return verifyKey, nil
	//})
	//fatal(err)
	//
	//claims := token.Claims.(*CustomClaimsExample)
	//fmt.Println(claims.CustomerInfo.Name)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run("x.tar:3000") // 监听并在 0.0.0.0:8080 上启动服务
}
