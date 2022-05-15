package main

import "fmt"

//传递int类型指针变量
func changeValue(p *int) {
	*p = 10
}

func main() {
	var a int = 20
	changeValue(&a)
	fmt.Println(a)
}
