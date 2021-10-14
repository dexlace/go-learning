package main

import "fmt"

// 全局变量
var x = 100

func f1(){
	fmt.Println(x)
}
func f2()  {
	// 局部作用域
	x:=99
	fmt.Println(x)
}

func main() {
	f1()
	f2()
	// 块作用域
	if x:=9;x<19 {
		fmt.Println(x)
	}
	fmt.Println(x)

}
