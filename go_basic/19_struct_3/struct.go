package main

import (
	"fmt"
)

// 1.结构体的字段类型
// 结构体的字段类型可以是 基本数据类型 切片 map 结构体
// 指针 slice和map的零值都是nil 即还没有分配空间
// 如果需要使用这样的字段，需要先make才能使用

type Person struct {
	name     string
	age      int
	hobby    []string
	mapValue map[string]string
}

// 2. 结构体嵌套

// 用户结构体

type User struct {
	userName string
	password string
	sex      string
	age      int
	address  Address
}

// 收货地址结构体

type Address struct {
	name  string
	phone string
	city  string
}

// 结构体的继承

type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动 \n", a.name)
}

type Dog struct {
	age int
	// 通过结构体嵌套完成继承
	Animal
}

func (d Dog) wang() {
	fmt.Printf("%v在汪汪汪\n", d.name)

}

func main() {

	var person = Person{}
	person.name = "张三"
	person.age = 12

	// 给切片申请内存空间
	person.hobby = make([]string, 4, 4)
	person.hobby[0] = "睡觉"
	person.hobby[1] = "吃饭"

	// 给map申请空间
	person.mapValue = make(map[string]string)
	person.mapValue["address"] = "深圳"
	person.mapValue["phone"] = "1234444"

	// 加入#打印完整信息
	fmt.Printf("%#v\n", person)

	fmt.Println("#####################结构体的继承##############")
	dog := Dog{
		age: 10,
		Animal: Animal{
			name: "hhhh",
		},
	}

	// 运行后，发现Dog拥有了父类的方法
	dog.run()
	dog.wang()

}
