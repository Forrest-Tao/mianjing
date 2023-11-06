package main

import "fmt"

func main() {
	var a int8 = -1
	var b int8 = -128 / a
	// -128为无类型常量，无类型常量除以变量得变量，-128/-1得128，128为int8的类型，溢出，得-128
	fmt.Println(b) //-128
}
