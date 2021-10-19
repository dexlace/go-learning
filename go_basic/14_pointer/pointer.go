package main

import "fmt"

// 这里的指针和c语言的指针一模一样
func main() {

	var a = 20

	// 得到a 的地址 &取的是a的地址
	// b这里是指针类型
	var b = &a
	// 打印的是a的地址
	fmt.Println(b)
	// 需要打印其值的话 用*b,即*指针的形式
	fmt.Println(*b)

	*b = 30
	// 改变b的值则会影响到a
	fmt.Println(a)
	/**
	对于引用类型的变量，我们在使用的时候不仅要声明它，
	还要为它分配内存空间，否则我们的值就没办法存储
	*/
	// 指针的new
	//var d *int
	//*d=100
	//fmt.Println(d)

	/**
		而new用于类型的内存分配
	    适用于对指针进行分配空间
	    如上，指针d的声明同时没有分配内存空间 会panic
	*/
	d := new(int)
	*d = 100
	fmt.Println(*d)

	// 而make用来分配的是slice map和channel的初始化
	// 详细见之前的相应demo
}
