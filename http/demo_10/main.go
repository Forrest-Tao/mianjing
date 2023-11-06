package main

import (
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func middleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 执行handler之前的逻辑
		next.ServeHTTP(w, r)
		// 执行handler之后的逻辑
	})
}

// logginHandler 中间件相当于中间包了一层，在next 前面写next之前的需要干的事  在next 后面写next之后的需要干的事
func logginHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		st := time.Now()
		log.Printf("start %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("comleted %s in %v", r.URL.Path, time.Since(st))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", logginHandler(http.HandlerFunc(index)))
	log.Println("Listening ...")
	http.ListenAndServe(":9090", mux)
}
