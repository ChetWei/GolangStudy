package main

import (
	"fmt"
	"time"
)

//有缓冲的channel

func main() {
	//创建一个缓冲为3的通道
	c := make(chan int, 3)
	//放入的元素数量,开辟大小
	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		fmt.Println("子go程运行")
		//当缓冲队列中的容量满了之后，会进入阻塞，等待被缓冲区内容被取走，再放入
		for i := 0; i < 4; i++ {
			c <- i + 1
			fmt.Println("发送元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	//先睡两秒，让子go程发送完成
	time.Sleep(time.Second * 2)

	//当channel为空，从里面取数据也会阻塞
	for i := 0; i < 4; i++ {
		//从c中接受并赋值
		num := <-c
		fmt.Println("接受 num=", num)
	}

	fmt.Println("main 结束")
}
