package lib1

import "fmt"

//对外开放函数首字母要大写
func B() {
	fmt.Println("lib12test...")
}
func init() {
	fmt.Println("lib12 init")
}
