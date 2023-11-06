package main

import "fmt"

func main() {
	str := "Hello, 世界"
	runes := []rune(str)
	for i, s := range str {
		fmt.Printf("i: %d %c\n", i, s)
	}
	fmt.Println("----------------")
	for i, r := range runes {
		fmt.Printf("Character %d: %c\n", i, r)
	}

}
