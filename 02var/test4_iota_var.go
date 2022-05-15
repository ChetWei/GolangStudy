package main

import "fmt"

/*
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/

//const (
//	a = iota //0
//	b        //iota +=1  1
//	c        //iota +=1  2
//	d = "ha" //独立值，iota += 1   3
//	e        //"ha"   iota += 1  4
//	f = 100  //iota +=1 	5
//	g        //100  iota +=1 6
//	h = iota //7, iota +=1 恢复计数
//	i        //8, iota +=1
//)

const (
	k = 1 << iota //左移 0 位，不变仍为 1
	l = 3 << iota //3 (011)  左移 1 位，变为二进制 110，即 6。
	m             //6 (110)左移2位，变为 12(1100)
	n             //12(1100)左移3位，变为 24(11000)
)

func main() {
	fmt.Println("k", k)
	fmt.Println("l", l)
	fmt.Println("m", m)
	fmt.Println("n", n)
}
