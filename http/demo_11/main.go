package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("D:/")))

	http.ListenAndServe(":9090", nil)
}
