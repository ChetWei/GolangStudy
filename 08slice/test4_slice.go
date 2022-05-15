package main

import "fmt"

func main() {
	var slice1 []int

	//空切片
	if slice1 == nil {
		fmt.Println("len slice ", len(slice1)) //没有分配slice空间 长度为0
	}

	slice1 = make([]int, 3) //开辟3个空间，int默认值为0
	fmt.Println("len slice ", len(slice1))
}
