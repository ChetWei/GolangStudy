package main

/*
引起panic(恐慌)的几个例子
空指针取值、下标越界、向nil map中添加键值对

*/

func main() {
	//1.空指针取值
	//var intPtr *int      //定义一个空指针
	//fmt.Println(*intPtr) // panic: runtime error: invalid memory address or nil pointer dereference

	//2.越界取值
	//mySlice := make([]int, 0)
	//mySlice = append(mySlice, 1, 2, 3, 4, 5, 6)
	//fmt.Println(mySlice[7]) //panic: runtime error: index out of range [7] with length 6

	//3.向空map中添加键值对
	var mmap map[string]int //这个map还没有分配内存空间
	mmap["name"] = 123      //panic: assignment to entry in nil map

}
