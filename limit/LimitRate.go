package limit

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func LimitRate() {
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

func LimitReal() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("OK!!!"))
	})
	_ = http.ListenAndServe(":8888", MyLimit(mux))
}
