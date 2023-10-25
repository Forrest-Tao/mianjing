package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

// 自定义Server 和Servermux
func main() {
	server := http.Server{
		Addr:         ":9090",
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", myFunc)
	// 其中func ListenAndServe(addr string, handler Handler) error 的handler有实现Handler
	//http.ListenAndServe(":9090", mux)
	server.Handler = mux
	server.ListenAndServe()
}
