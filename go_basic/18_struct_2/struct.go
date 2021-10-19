package main

import "fmt"

// 注意方法的概念 他是函数  但是是特定的一种作用于特定类型变量的函数
// 基于任意类型定义其方法都行
// 例如定义在自己定义类型上的方法
type myInt int

func fun(x int, y int) int {
	return x + y
}

func (m myInt) PrintMyIntInfo() {
	fmt.Println("我是自定义类型里面的自定义方法")

}

// 定义于结构体的方法即可视作java类中的函数
//////结构体方法和接收者

// 先定义一个结构体

type Person struct {
	name string
	age  int
	sex  string
}

//在go语言中，没有类的概念
//但是可以给类型（结构体，自定义类型）定义方法。
//所谓方法就是定义了接收者的函数。
//接收者的概念就类似于其他语言中的this 或者self

// func (接收者变量 接收者类型) 方法名(参数列表)(返回参数) {
//    函数体
//}

// 接收者中的参数变量
// 官方建议使用接收者类型名的第一个小写字母
//例如，Person类型的接收者变量应该命名为p，
//Connector类型的接收者变量应该命名为c等。

// 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型
// 非指针类型：表示不修改结构体的内容
// 指针类型：表示修改结构体中的内容

func (p Person) PrintInfo() {
	fmt.Println("姓名： ", p.name)
	fmt.Println("年龄： ", p.age)
	fmt.Println("性别： ", p.sex)
}

// 因为结构体是值类型 我们修改的时候，必须传入的指针

func (p *Person) SetInfo(name string, age int, sex string) {
	p.name = name
	p.age = age
	p.sex = sex
}

// 这里可以封装结构体的一种构造函数？类似于类中的构造函数
// 其实就是返回一个结构体类型的变量即可
// 也不是很。。。
func newPerson(name string) Person {
	return Person{
		name: name,
	}
}

// 结构体的匿名字段
// 结构体允许其成员字段在声明时没有字段名而只有类型
// 匿名字段默认采用类型名作为字段名
// 结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能一个

type Student struct {
	string
	int
}

func main() {
	p1 := Person{
		"张三",
		18,
		"male",
	}
	p1.PrintInfo()
	fmt.Println("**********这里的set方法是引用传递***************")
	p1.SetInfo("李四", 29, "female")
	p1.PrintInfo()

	// 构造方法

}
