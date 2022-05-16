package main

import "fmt"

func main() {
	var a string
	//a的内部 pair<statictype:string,value:"hello>
	a = "hello"

	var allType interface{}
	allType = a //allType的内部 pair<type:string,value:"hello>

	value, _ := allType.(string)
	fmt.Println(value)
}
