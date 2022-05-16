package main

import (
	"fmt"
	"reflect"
)

//结构体标签
type resume struct {
	//添加字段标签
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()

	for i := 0; i < t.NumField(); i++ {
		//获取tag里面的 info对应的值
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, "| doc: ", tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
