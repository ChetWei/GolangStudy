package main

import "fmt"

/*
恐慌，一种严厉的异常
*/

func caculate(a int) int {
	if a == 0 {
		panic("除数不能为负数") //主动判断，抛出恐慌
	}
	return 100 / a //如果没有处理，a=0将会引起恐慌
}

//a 为负数的情况，出现恐慌
//panic: runtime error: integer divide by zero

func main() {

	//在引起恐慌之前会执行这里
	defer func() {
		fmt.Println("defer")
		//recover从恐慌中复活，找到导致恐慌的原因
		if err := recover(); err != nil {
			fmt.Println("恐慌原因", err)
		}
	}()

	//将0作为除数，引起恐慌
	result := caculate(0)

	//如果上面恐慌，这里的defer将无法执行
	defer func() {
		fmt.Println("defer2")
	}()

	fmt.Println(result)
}
