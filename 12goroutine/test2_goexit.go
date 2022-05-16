package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//go协程之间的执行是无序的

	//使用go协程 定义并调用匿名函数 形参为空，返回值为空
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前的go协程
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//带有形参的go匿名函数
	go func(a int, b int) bool {
		fmt.Println("a=", a, ",b=", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
