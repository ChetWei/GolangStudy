package main

// const 定义的变量不能修改

import (
	"fmt"
	"unsafe"
)

//常量枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH

	sizea := unsafe.Sizeof(a)
	fmt.Println("sizea: ", sizea)

	fmt.Println("面积为 : %d", area)
	println(a, b, c)
	fmt.Println("unknown=", Unknown)
}
