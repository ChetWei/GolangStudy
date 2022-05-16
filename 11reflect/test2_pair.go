package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

//实现Reader接口
func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

//实现Write接口
func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	//b: pair<type:Book,value:book{}的地址>
	b := &Book{}

	//声明一个 Reader接口变量
	//r: pair<type,value> 暂时为空
	var r Reader
	//r: pair<type:Book,value:book{}的地址>
	r = b //指向Book具体类，因为Book实现了Reader的接口
	r.ReadBook()

	//声明一个 Writer接口变量
	var w Writer
	//r: pair<type:Book,value:book{}的地址>
	w = r.(Writer) //这里对r 进行断言为Writer，因为r和w具体的type是一致的都是Book
	w.WriteBook()

}
