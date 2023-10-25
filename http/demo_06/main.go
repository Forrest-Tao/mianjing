package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	//自定义 http.Server
	server := http.Server{
		Addr:         ":9090",
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	http.HandleFunc("/", myFunc)
	server.ListenAndServe()
}
