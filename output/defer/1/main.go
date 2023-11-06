package main

import "fmt"

type A struct {
}

func (a A) f(i int) A {
	fmt.Println(i)
	return a
}
func main() {

	a := A{}
	defer a.f(1).f(2).f(3).f(4)

	fmt.Println("done")
}
