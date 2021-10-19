package main

import "fmt"

// 闭包
//闭包可以理解成 “定义在一个函数内部的函数”定义在一个函数内部的函数
// 可以让一个变量不污染全局
// 可以让一个变量常驻内存

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 闭包就是将函数内部 和 函数外部连接起来的桥梁
// 闭包想做的是把f2塞进f1的参数里
// 即可以将f2的对应的函数类型作为参数
// f1()对应的参数类型作为返回值
func f3(f func(x, y int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}

// 又比如
// 闭包的写法：函数里面嵌套一个函数，最后返回里面的函数就形成了闭包
// 闭包=函数+外部变量的引用
func adder() func() int {
	i := 10
	return func() int {
		return i + 1
	}
}

//注意：由于闭包里作用域返回的局部变量资源不会被立刻销毁，
//所以可能会占用更多的内存，
//过度使用闭包会导致性能下降，建议在非常有必要的时候才使用闭包。

func main() {
	// 则能调用
	f1(f3(f2, 120, 119))

	fadd := adder()
	fmt.Println(fadd())
	fmt.Println(fadd())
	fmt.Println(fadd())

}
