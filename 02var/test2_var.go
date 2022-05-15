package main

import "fmt"

func main() {
	var a int = 9
	var b = a

	fmt.Println("a addr:", &a)
	fmt.Println("b addr:", &b)

}
