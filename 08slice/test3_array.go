package main

import "fmt"

/*
切片的引用传递
*/

func printArray2(myArray []int) {
	//使用切片动态数组参数的时候，使用的是引用传递
	fmt.Println("----------------printArray----------")
	for _, value := range myArray {
		fmt.Println(" value= ", value)
	}
	myArray[0] = 1000
}

func main() {

	//动态数组，切片slice，根据元素个数自动分配空间
	var myArray = []int{10, 20, 30}

	//切片引用传递
	printArray2(myArray)

	printArray2(myArray)
}
