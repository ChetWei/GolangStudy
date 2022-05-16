package main

import "fmt"

//定义一个接口，本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色

}

//具体类1
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep...")
}

func (this *Cat) GetColor() string {
	return this.color
}

//具体类2
type Dog struct {
	color string //猫的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep...")
}

func (this *Dog) GetColor() string {
	return this.color
}

//多态
func showAnimal(animal AnimalIF) {
	//传递进来的是引用
	animal.Sleep() //多态，会根据类的具体类型，调用具体的实现方法
	fmt.Println("color=", animal.GetColor())
}

func main() {
	//1 >>> 直接通过接口指针进行具体类的调用
	/*	var animal AnimalIF //接口数据类型，父类指针

		animal = &Cat{"Green"}
		animal.Sleep()

		animal = &Dog{"yellow"}
		animal.Sleep()*/

	//2 >>> 通过多态的方式进行具体类的调用
	cat := Cat{"green"}
	dog := Dog{"yellow"}
	showAnimal(&cat) //传递地址进去
	showAnimal(&dog)
}
