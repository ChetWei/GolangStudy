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

	for {
		//ok如果为true表示channel没有关闭,否则关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			//通道关闭退出循环
			break
		}
	}
	fmt.Println("main finished ..")
}
