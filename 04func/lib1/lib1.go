package lib1 //包名与所在目录名相同
import "fmt"

//对外开放函数首字母要大写
func A() {
	fmt.Println("lib1test...")
}

func init() {
	fmt.Println("lib1 init")
}
