package main

import "fmt"

func main() {
	const a int8 = -1
	var b int8 = -128 / a
	// -128 / a (constant 128 of type int8) overflows int8
	fmt.Println(b) //panic
}
