package main

import (
	"fmt"
)

/*
用户自定义错误
*/
type CustomError struct {
	Msg string //错误信息
}

func (this *CustomError) Error() string {
	return "自定义异常:" + this.Msg
}

func caculate3(a int) (result int, err error) {
	if a == 0 {
		//panic抛出自己定义的错误,但是还会导致后面的程序中止
		panic(&CustomError{"除数不能为0"})
	}

	//不在范围之内，温和的返回错误
	if a < 5 || a > 50 {
		err = &CustomError{"不在范围之内[5,50]"}
		return //错误直接返回
	}

	//正常情况，错误为nil
	return 100 / a, nil
}

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
	res, err := caculate3(0)
	if err != nil {
		fmt.Println("错误信息", err)
	}

	defer func() {
		fmt.Println("defer2")
	}()

	fmt.Println("结果为: ", res)
}
