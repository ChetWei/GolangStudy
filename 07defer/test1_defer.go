package main

import "fmt"

func main() {

	//defer关键字，类似c++的析构函数，在函数完成后执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") //有多个defer的时候，defer的执行是栈的结构，出栈执行

	fmt.Println("main hello go 1")
	fmt.Println("main hello go 2")
}
