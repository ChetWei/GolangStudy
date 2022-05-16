package main

import "fmt"

//将channel变量作为参数
func fibonacii(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		//使用select
		select {
		case c <- x:
			//如果检测到c可写，则这个case就会进来
			t := x
			x = y
			y = t + y

		case data := <-quit:
			//检测到quit可读
			fmt.Println("quit", data)
			return
		}
	}
}

func main() {
	//创建两个channel
	c := make(chan int)
	quit := make(chan int)

	//子程
	go func() {
		for i := 0; i < 10; i++ {
			//读取
			fmt.Println(<-c)
		}
		quit <- 0 //写入
	}()

	//main
	fibonacii(c, quit)
}
