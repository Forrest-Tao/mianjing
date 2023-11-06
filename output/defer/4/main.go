package main

import "fmt"

// inner if segment done  101
func main() {
	n := 1
	if n == 1 {
		defer func() {
			fmt.Println(n)
		}()
		fmt.Println("inner if segment")
	}
	n += 100
	fmt.Println("done")
}
