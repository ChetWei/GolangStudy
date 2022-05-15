package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func call")
	return 0
}

func returnFunc() int {
	fmt.Println("rerutn func call")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	//当有return和defer的时候，先执行return 再执行 defer
	returnAndDefer()
}
