package main

import "fmt"

func main()  {

	fmt.Println("/*************************GoLang的变量****************************/")
	fmt.Println("hello world")
	fmt.Println("A", "B", "C")
	fmt.Println("C:\\Users\\Administrator\\Desktop")
	fmt.Println("天龙八部\r你好")


	// 变量以var开头
	var name = "zhangsan"
	fmt.Println(name)
	// 同时定义类型
	var name2 string = "zhangsan2"
	fmt.Println(name2)
	// 类型一样的多个变量
	var a1, a2 string
	a1 = "lisi"
	a2 = "wangwu"
	fmt.Println(a1)
	fmt.Println(a2)

	// 格式化输出
	fmt.Printf("name=%v name2=%v \n", name, name2)

	// 定义常量
	const pi = 3.14
	const (
		A = "A"
		B = "B"
	)

	// Const常量结合iota的使用
	// iota是golang 语言的常量计数器，只能在常量的表达式中使用
	// iota在const关键字出现时将被重置为0
	const a = iota
	const (
		b = iota // b=0
		c        // c=1
		d        // d=2
	)

	const (
		e = iota //e=0
		_        // 跳过某些值
		f        //f=2
	)

	var number = 15
	fmt.Printf("原样输出 %v\n", number)
	fmt.Printf("十进制输出 %d\n", number)
	fmt.Printf("二进制输出 %b\n", number)
	fmt.Printf("八进制输出 %o\n", number)
	fmt.Printf("十六进制输出 %x\n", number)
}