package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myFunc)
	// 其中func ListenAndServe(addr string, handler Handler) error 的handler有实现Handler
	http.ListenAndServe(":9090", mux)
}
