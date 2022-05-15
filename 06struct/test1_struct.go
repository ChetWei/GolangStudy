package main

import "fmt"

//给类型取别名
type myint int

//定义结构体
type Book struct {
	title string
	price float32
}

func changeBook(book Book) {
	//传递的是book的一个副本
	book.price *= 100
}

func changeBook2(book *Book) {
	//传递的是book的一个指针
	book.price *= 100
}

func main() {

	var book1 Book //声明一个Book类型
	book1.title = "Golang"
	book1.price = 12.34

	fmt.Println(book1)
	changeBook(book1)
	fmt.Println(book1)
	changeBook2(&book1)
	fmt.Println(book1)
}
