package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//类里面的方法
func (this *User) Call() {

	fmt.Println("user is called ..")
	fmt.Printf("%v \n", this)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType)
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	inpuNumField := inputType.NumField() //类里的字段数量
	fmt.Println("inpuNumField is:", inpuNumField)

	//通过type获取里面的字段
	for i := 0; i < inpuNumField; i++ {
		field := inputType.Field(i)              //获取这个字段的信息
		value := inputValue.Field(i).Interface() //获取这个字段的具体实现值
		//字段的命名 field.Name
		//字段类型 field.Type
		fmt.Printf("%s : %v = %v \n", field.Name, field.Type, value)
	}
}

func main() {
	user := User{1, "Tom", 18}

	DoFiledAndMethod(user)
}
