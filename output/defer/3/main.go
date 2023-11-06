package main

import "fmt"

// done  101
func main() {
	n := 1
	defer func() {
		fmt.Println(n)
	}()
	n += 100
	fmt.Println("done")
}
