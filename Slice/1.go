package main

import (
	"fmt"
)

func main() {
	var array [20]int
	var slice = array[5:6]
	fmt.Println("lenth of slice: ", len(slice))
	/*切片的容量（cap）表示的是在底层数组中从切片的起始元素（slice[0]）到数组末尾元素之间的元素数量*/
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
	/*	fmt.Println(array)
		fmt.Println(slice)*/
}
