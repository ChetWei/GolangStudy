package main

import (
	"fmt"
	"time"
)

func task1() {
	i := 0
	for {
		i++
		fmt.Println("task1 : i=", i)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	//创建一个go协程，去执行task1的方法
	go task1()

	//主进程
	i := 0
	for {
		i++
		fmt.Println("main task : i=", i)
		time.Sleep(1 * time.Second)
	}
}
