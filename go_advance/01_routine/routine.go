package main

import (
	"fmt"
	"time"
)

// goroutine：用户态的线程
// goroutine是由go的运行时（runtime）调度和管理的
// go程序会智能的将goroutine中的任务合理的分配给每个CPU
// go语言在语言层面上已经内置了调度和上下文切换的机制

func hello() {
	fmt.Println("hello")

}
func main() {
	// 1.启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个`go`关键字
	// 开启一个单独的goroutine去执行hello函数
	go hello()
	fmt.Println("main")

	// 执行结果会打印main，而不会显示hello
	// 这是因为main线程结束了 栈空间被销毁了  go hello没执行完只能时销毁了
	// 可以在这里等一秒试试
	time.Sleep(time.Second)

	// 2.闭包从外面拿参数产生的并发问题
	// 以下会出现打印多个相同数字的问题
	for i := 1; i < 1000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	// 解决 传参进来
	for i := 1; i < 1000; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// 3. goroutine什么时候结束，goroutine对应的函数结束了，goroutine结束了
	// main函数执行完了 由main函数创建的goroutine也结束了
}
