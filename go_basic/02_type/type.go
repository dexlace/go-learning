package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main(){
	var num1 = 12
	fmt.Println(unsafe.Sizeof(num1))

	// 类型转换
	var b1 int16 = 10
	var b2 int32 = 12
	var b3 = b2 + int32(b1)
	fmt.Println(b3)

	// 高位转低位时可能会精度丢失
	// 比如16位转8位
	var n1 int16 = 130
	fmt.Println(int8(n1))

	fmt.Println("/*************************GoLang的数据类型 注意类型推导方式定义变量 以字面量为例****************************/")

	// 变量名 := 表达式
	// 字面量语法
	// 二进制的1111
	i1 := 0b1111

	// 八进制的1111
	i2 := 01111

	i3 := 0x1111

	fmt.Printf("%d\n", i1)
	fmt.Printf("%d\n", i2)
	fmt.Printf("%d\n", i3)

	// 浮点类型
	var pai = math.Pi

	// 打印浮点类型 默认小数点后6位
	fmt.Printf("%f\n", pai)
	// 打印浮点类型 小数点后2位
	fmt.Printf("%.2f\n", pai)

	// 布尔类型
	var f1 = false
	if f1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
