package main

import "fmt"

func printArray(myArray [3]int) {
	//有指定长度数组时，使用的是值拷贝
	fmt.Println("----------------printArray----------")
	for index, value := range myArray {
		fmt.Println(" index= ", index, " value= ", value)
	}
	myArray[0] = 1000
}

func main() {

	var myArray = [3]int{10, 20, 30}

	printArray(myArray)

	printArray(myArray)
}
