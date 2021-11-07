package main

import (
	"fmt"
	"sync"
)

// 互斥锁是传统并发编程中对共享资源进行访问控制的主要手段，
// 它由标准库sync中的Mutex结构体类型表示。
// sync.Mutex类型只有两个公开的指针方法，Lock和Unlock。
// Lock锁定当前的共享资源，Unlock 进行解锁
// 定义一个锁
//var mutex sync.Mutex
//// 加锁
//mutex.Lock()
//// 解锁
//mutex.Unlock()

var count=0
var wg sync.WaitGroup
var lock sync.Mutex

func test(){
	for i:=0;i<50000;i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
	wg.Done()
}


func main() {
	wg.Add(2)
	go test()
	go test()

	wg.Wait()
	fmt.Println(count)

}
