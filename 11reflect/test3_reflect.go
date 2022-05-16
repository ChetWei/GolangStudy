package main

import (
	"fmt"
	"reflect"
)

//通过反射获取参数的数据类型
func reflectNum(arg interface{}) {

	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345

	reflectNum(num)
}
