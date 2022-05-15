package main

import "fmt"

// 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数
//返回类型为int
func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

//多返回值,匿名返回两个int值
func foo2(a string) (int, int) {
	return 1, 2
}

//多返回值,有形参名称
func foo3(a string) (r1 int, r2 int) {
	r1 = 3 //对要返回的形参进行赋值,没有赋值之前是类型默认值
	r2 = 4
	// return 3,4
	return //返回形参
}

//多返回值,有形参名称,多个返回类型相同的情况
func foo4(a string) (r1, r2 int) {
	return 3, 4
}

func main() {
	c := foo1("abc", 600)

	fmt.Println("c=", c)

	r1, r2 := foo2("hello")
	fmt.Println(r1, r2)

	r3, r4 := foo3("hi")
	fmt.Println(r3, r4)
}
