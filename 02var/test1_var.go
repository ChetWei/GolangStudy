package main

import "fmt"

/*
变量的声明与初始化

*/

//定义的变量，必须要使用，否则会编译错误

//声明全局变量 方法一、二、三 是没有问题的

func main() {
	//方法一：声明变量,默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T \n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T \n", b)

	//多变量声明
	var bb, bc int = 200, 300
	fmt.Println("bb = ", bb)
	fmt.Println("bc = ", bc)

	//多行变量的声明
	var (
		bd int     = 200
		be float32 = 900.8
	)
	fmt.Println(bd, be)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T \n", c)

	var d = "hello"
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T \n", d)

	//方法四 使用  :=  省去var 关键字 和 变量类型 ！只能在函数体内使用
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T \n", e)

	f := 100.01
	fmt.Println("e = ", f)
	fmt.Printf("type of f = %T \n", f)
}
