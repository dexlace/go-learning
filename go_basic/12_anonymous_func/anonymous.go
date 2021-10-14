package main

import "fmt"

// 匿名函数即没有名字的函数
// func(x,y int){
//	fmt.Println(x+y)
//}即为匿名函数
// f1被赋值为一个函数类型  注意函数类型的概念
// 即拿到匿名函数的地址了
var f1=func(x,y int){
	fmt.Println(x+y)
}
// 但是上述匿名函数的使用不太好，毕竟匿名函数一般使用在函数内部
func main() {
	f1(5,6)

	// 函数内部定义匿名函数
	// f2被赋值为一个函数类型  注意函数类型的概念
	f2:= func(x,y int) {
        fmt.Println(x+y)
	}

	f2(10,20)
	f2(20,30)

	// 如果匿名函数只执行一次，则可简写为立即执行函数
	func(x,y,z int){
		fmt.Println(x+y+z)
	}(1,2,3)

	func(m,n int){
		fmt.Println(m-n)
	}(1,2)

}
