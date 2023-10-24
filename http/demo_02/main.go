package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi")
}

func main() {
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}
