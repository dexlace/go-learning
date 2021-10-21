package main

import "fmt"

// 1.结构体实现多个接口
// 也就是实现多个接口的所有方法

type Animal1 interface {
	SetName(string)
}

type Animal2 interface {
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

// 2. 接口嵌套
// 允许接口嵌套接口，
// 我们首先创建一个 Animal1 和 Animal2 接口，
// 然后使用Animal接受刚刚的两个接口，实现接口的嵌套

type Animal11 interface {
	SetName(string)
}

type Animal22 interface {
	GetName() string
}

type Animal interface {
	Animal1
	Animal2
}

type Dog2 struct {
	Name string
}

func (d *Dog2) SetName(name string) {
	d.Name = name
}
func (d Dog2) GetName() string {
	return d.Name
}

func main() {

	// 1.结构体实现多个接口
	var dog = Dog{
		"小黑",
	}
	dog.SetName("小白")
	fmt.Println("狗名：" + dog.GetName())

	// 2.接口嵌套
	// 注意取地址才能赋值给Animal
	var dog2 = &Dog2{
		"大黄",
	}
	var d Animal = dog2
	d.SetName("大橙")
	fmt.Println(d.GetName())
}
