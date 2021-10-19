package main

import "fmt"

//type 接口名 interface {
//	方法名1 (参数列表1) 返回值列表1
//	方法名2 (参数列表2) 返回值列表2
//}

// 使用type将接口定义为自定义的类型名
// Go语言的接口在命名时，一般会在单词后面添加er
// 如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等，接口名最好突出该接口的类型含义

// 当方法名首字母是大写且这个接口类型名首字母也是大写时，
//这个方法可以被接口所在的包（package）之外的代码访问

//参数列表和返回值列表中的参数变量名是可以省略
// 接口是一种类型

type Speaker interface {
	speak()
}

type Dog struct {
}

// dog结构体的方法
func (d Dog) speak() {
	fmt.Println("汪汪汪")
}

type Cat struct {
}

// cat结构体的方法
func (c Cat) speak() {
	fmt.Println("喵喵喵")
}

// 以上两个结构体都实现了speak()方法，那么就可以作为结构体Speaker的一种类型

// 下面定义一个函数
func da(s Speaker) {
	s.speak()
}
func main() {
	da(new(Dog))
	da(new(Cat))

}
