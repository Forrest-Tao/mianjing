package main

import (
	"fmt"
)

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

/*考察Slice的扩容，*/
/*
如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍；
如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍；

append函数执行时会判断切片容量是否能够存放新增元素，如果不能，
则会重新申请存储空间，新存储空间将是原来的2倍或1.25倍（取决于扩展原空间大小），
本例中实际执行了两次append操作，第一次空间增长到4，所以第二次append不会再扩容，
所以新旧两个切片将共用一块存储空间。程序会输出”true”
*/
func main() {
	var slice []int
	slice = append(slice, 1, 2, 3)
	newSlice := AddElement(slice, 4)
	fmt.Println(&slice[0] == &newSlice[0])

}
