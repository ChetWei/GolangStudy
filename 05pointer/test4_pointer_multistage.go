package main

import "fmt"

//测试多级指针

func main() {

	var a int = 10

	//指针变量，指向a的地址
	var pt *int = &a

	//二级指针
	var pp **int
	pp = &pt //存放pt指针变量的地址

	fmt.Println(&pt) //pt指针变量的地址
	fmt.Println(pp)

	fmt.Println(*pt)  //取pt存放地址指向的值
	fmt.Println(*pp)  // *pp 得到的是 pt指针变量
	fmt.Println(**pp) // **pp 得到的是 pt指针变量指向的值

}
