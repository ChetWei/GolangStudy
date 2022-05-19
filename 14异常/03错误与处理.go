package main

import (
	"errors"
	"fmt"
)

/*
错误，一种比较温和的异常
*/

func caculate2(a int) (result int, err error) {
	if a == 0 {
		panic("恐慌，除数不能为负数") //主动判断，抛出恐慌
	}

	//不在范围之内，温和的返回错误
	if a < 5 || a > 50 {
		err = errors.New("错误，a的输入范围[5,50]")
		return //错误直接返回
	}

	//正常情况，错误为nil
	return 100 / a, nil
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

	//这里引起error，error不会导致后面的程序中断
	res, err := caculate2(3)
	if err != nil {
		fmt.Println("错误信息", err)
	}

	defer func() {
		fmt.Println("defer2")
	}()

	fmt.Println("结果为: ", res)
}
