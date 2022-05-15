package main

import "fmt"

/*
切片的相关操作
*/

func main() {
	//创建一个切片，最大容量 cap=5，3个默认值0
	var numbers = make([]int, 3, 5)
	fmt.Println("numbers", numbers)
	fmt.Printf("numbers len = %d, cap=%d,slice=%v \n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	numbers = append(numbers, 2)

	//切片扩容机制
	//此时cap的容量已经满了，会动态开辟新的容量，开辟的容量为 +cap
	numbers = append(numbers, 3)
	numbers = append(numbers, 4)
	fmt.Printf("numbers len = %d, cap=%d,slice=%v \n", len(numbers), cap(numbers), numbers)
}
