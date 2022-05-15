package main

import "fmt"

func main() {
	//声明数组
	var blance [10]float32

	var i int
	for i = 0; i < 10; i++ {
		blance[i] = float32(i * 2.0)
	}

	for i = 0; i < 10; i++ {
		fmt.Println(blance[i])
	}
	//声明并初始化数组
	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//通过字面量在声明数组的同时快速初始化
	balance3 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//如果数组的长度不确定，可以使用...代替数组的长度，编译器会根据元素个数自行推断数组的长度
	balance4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	fmt.Println(balance2)
	fmt.Println(balance3)
	fmt.Println(balance4)
}
