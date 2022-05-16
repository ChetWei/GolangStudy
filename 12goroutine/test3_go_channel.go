package main

import "fmt"

//goroutine之间的通信 channel

func main() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine函数")
		fmt.Println("goroutine 运行")
		//发送数据到channel
		c <- 666
	}()
	// c 进行阻塞，等待取数据
	//从c中接受数据，赋值给num
	num := <-c
	fmt.Println("num=", num)

}
