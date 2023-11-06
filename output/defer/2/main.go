package main

import "fmt"

func f(n int) {
	defer fmt.Println(n)
	n += 100
}

// 1 2 done
func main() {

	f(1)
	f(2)
	fmt.Println("done")
}
