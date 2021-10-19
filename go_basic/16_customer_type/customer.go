package main

import "fmt"

// 自定义类型
// 即定义了一种类型而已
type myint int

// 自定义类型别名
// 仅仅是别名而已 类型还是原类型
type yourInt = int

func main() {
	var n myint
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 1000
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	fmt.Println("rune是int32的类型别名")
	// 所以rune是int32的类型别名
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%c\n", c)
	fmt.Printf("%T\n", c)
}
