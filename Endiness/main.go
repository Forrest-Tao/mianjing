package main

import (
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool {
	var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01
	// 大端(16进制)：00 00 00 0
	// 小端(16进制)：01 00 00 00
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)
	if *pb != 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsLittleEndian())
}

// 运行结果：ture
