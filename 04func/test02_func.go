package main

import (
	"04func/lib1"
	mylib2 "04func/lib2" //给包起别名，可以使用.别名，这样调用的时候就可以不使用名称来调用，而是直接调用方法
)

func main() {
	lib1.A()
	lib1.B()

	//使用包别名调用方法
	mylib2.C()
}
