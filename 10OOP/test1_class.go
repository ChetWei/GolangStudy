package main

import "fmt"

/*
类名、属性名、方法名
首字母大写表示对外（其他包可以访问，否则只能够在本包内访问）
*/

//定义一个Hero类
type Hero struct {
	Name  string
	Ad    int
	Level int
}

//给结构体绑定方法
func (this Hero) showName() {
	//this 是调用该方法的对象的一个副本（拷贝）
	fmt.Println("Name =", this.Name)
}

func (this *Hero) GetName() string {
	//this 是调用该方法的对象的一个引用
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{"zhangsan", 100, 1}

	hero.showName()

	hero.SetName("lisi")

	hero.showName()

	name := hero.GetName()
	fmt.Println(name)

}
