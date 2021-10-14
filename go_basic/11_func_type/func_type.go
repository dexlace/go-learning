package main

import (
	"fmt"
)

func f1(){
	fmt.Println("f1")
}

func f2() int{
	return 5
}

// 函数作为参数类型
// 由于f2恰好是func() int类型  所以可以传进去
func f3(x func() int){
	ret:=x()
	fmt.Println(ret)

}

func f4(x,y int) int  {
	return x+y
}
func f5(x func(m,n int) int) int {
	// 注意要给作为参数的函数的参数进行赋值
	m:=5
	n:=6
	ret:=x(m,n)
	return ret

}



func f6(a,b int) int  {
	return a+b

}

// 函数作为返回值
func f7() func(int,int) int  {
	return f6
}



func main() {

	fmt.Println("-----------------函数类型------------------------")
	// 此时a为func ()类型
	// 即没有参数没有返回值的函数类型
	// 注意这里的赋值不是函数调用 没有带()
	a:=f1
	fmt.Printf("%T\n",a)
	b:=f2
	fmt.Printf("%T\n",b)
	// 打印出来的是地址
	fmt.Println(b)
	// 此时可以用b()进行函数调用
	fmt.Println(b())

	fmt.Println("----------函数作为参数类型---------")
	// 注意不要写成函数调用的形式
	f3(f2)
	fmt.Println(f5(f4))

	fmt.Println("----------函数作为返回值---------")
	// 相当于返回了一个函数的地址
	c:=f7()
	fmt.Println(c)
	fmt.Println(c(6,18))


}
