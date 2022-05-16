package main

import "fmt"

//万能的 数据类型

//interface{}是万能的数据类型 所有的基本数据类型都实现了interface
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called..")
	fmt.Println(arg)

	//使用 interface{} 如何区分 引用的底层数据类型是什么？

	// 给interface{} 提供"类型断 言"的机制
	value, ok := arg.(string) //如果是string ok=true
	if ok {
		fmt.Println("arg is string type,value= ", value)
	} else {
		fmt.Println("arg is not string type")
		fmt.Printf("value type is %T \n", value)
	}
}

type Book struct {
	title string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
