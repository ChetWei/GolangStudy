package main

import "fmt"

func main() {
	/*1.数组指针，一个指针，数据类型为数组（指向数组）*/
	//定义一个数组
	var arr = [3]int{10, 200, 3000}
	//定义数组指针
	var arrPt *[3]int
	//将数组的地址赋值给指针变量
	arrPt = &arr
	fmt.Println("arrpt=", arrPt) //arrpt= &[10 200 3000]
	fmt.Println("&arr=", &arr)   //输出内容等同于上面
	//输出地址的形式
	fmt.Printf("arr 数组的地址为：%p\n", &arr)
	fmt.Printf("arrPtr 存储的地址为：%p\n", arrPt)
	fmt.Println("*arrpt=", *arrPt) //输出的是数组的值

	//通过指针访问数组
	fmt.Println((*arrPt)[0]) //为了保证寻址运算符号* 优先于 [] ,使用括号

	/*2.指针数组，一个数组，里面的使用元素都为地址值*/
	a, b, c, d := 1, 2, 3, 4
	var ptrArr [4]*int               //声明一个指针数组，里面存放的都是指针
	ptrArr = [4]*int{&a, &b, &c, &d} //初始化
	fmt.Println(ptrArr)              //输出的是数组里面每个元素的地址
	fmt.Println(*ptrArr[0])          //取值

}
