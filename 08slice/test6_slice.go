package main

import "fmt"

//切片的截取，复制
func main() {
	s := []int{1, 2, 3}

	//截取操作，s1 指向的内存值和 s是一致的
	s1 := s[0:2]

	//如果s1修改了，那么s的值也会被修改
	s1[0] = 55

	fmt.Println(s)
	fmt.Println(s1)

	//copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
}
