package main

import "fmt"

//结构体类的继承
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

type SuperMan struct {
	Human  //SupeMan类继承了Human类的成员和方法
	height float32
}

//重新定义父类的方法Eat
func (this *SuperMan) Eat() {
	fmt.Println("Superman.Eat()...")
}

//子类的新方法
func (this *SuperMan) Walk() {
	fmt.Println("Superman.Walk")
}

func main() {
	human := Human{"zhangsan", "female"}
	human.Eat()

	superMan := SuperMan{human, 160}
	superMan2 := SuperMan{Human{"lisi", "male"}, 180}

	superMan.Eat()
	superMan.Walk()

	superMan2.Eat()
	superMan2.Walk()

}
