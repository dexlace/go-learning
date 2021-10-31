package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 0.怎么产生随机数
func f() {
	rand.Seed(time.Now().UnixNano())
	// 如果没有下面这行随机数的种子，那么每次运行都是一样的随机结果
	for i := 0; i < 3; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

// 1. 协程计数器

//WaitGroup 内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：
//
//Add() 用来添加计数
//Done() 用来在操作结束时调用，使计数减一，翻看源码可以看到，该方法的实现实际上就是调用 wg.Add(-1)
//Wait() 用来等待所有的操作结束，即计数变为 0，该函数会在计数不为 0 时等待，在计数为 0 时立即返回

var wg sync.WaitGroup

func main() {
	f()

	wg.Add(2) // 因为有两个goroutine，所以增加2个计数
	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // 操作完成，减少一个计数
	}()
	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // 操作完成，减少一个计数
	}()

	wg.Wait() // 等待，直到计数为0

	// 2. 设置Go并行运行的时候占用的cpu数量
	// runtime.GOMAXPROCS（）函数设置当前程序并发时占用的CPU逻辑核心数
	//  默认值是机器上的CPU核心数
	// 获取cpu个数
	npmCpu := runtime.NumCPU()
	fmt.Println("cup的个数:", npmCpu)
	// 设置允许使用的CPU数量
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

}
