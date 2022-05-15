package main

import "fmt"

/*
map是一种k-v数据类型，无序
*/

func main() {

	//=====> 第一种声明方式

	//声明myMap1是一个map类型，key是string，value是string
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map!")
	}

	//使用map前，需要使用make给map分配空间
	myMap1 = make(map[string]string, 10)

	//给map变量赋值
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// ==> 第二种声明方式
	myMap2 := make(map[string]int)
	myMap2["java"] = 1
	myMap2["c++"] = 2
	myMap2["python"] = 3
	fmt.Println(myMap2)

	// ==> 第三种声明方式
	myMap3 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap3)

}
