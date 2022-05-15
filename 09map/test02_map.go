package main

import (
	"fmt"
)

//map传参使用的是引用传递
func modifyMap(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历 无序的
	for key, value := range cityMap {
		fmt.Println("key =", key)
		fmt.Println("value = ", value)
	}

	//修改
	cityMap["USA"] = "DC"

	fmt.Println("cityMap: ", cityMap)

	//删除
	delete(cityMap, "China")

	//访问
	fmt.Println("cityMap[\"China\"]: ", cityMap["China"])

	modifyMap(cityMap)

	fmt.Println(cityMap)

}
