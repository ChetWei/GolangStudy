package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭channel
		close(c)
	}()

	//使用range迭代读取channel，当channel中没有数据时候等待，当channel关闭的时候退出
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finished")
}
