package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//  互斥锁的本质是当一个goroutine访问的时候，其他goroutine都不能访问。
// 这样在资源同步，避免竞争的同时也降低了程序的并发性能。程序由原来的并行执行变成了串行执行。
//
//其实，当我们对一个不会变化的数据只做“读”操作的话，是不存在资源竞争的问题的。
//因为数据是不变的，不管怎么读取，多少goroutine同时读取，都是可以的。
//
//所以问题不是出在“读”上，主要是修改，也就是“写”。修改的数据要同步，这样其他goroutine才可以感知到。
//所以真正的互斥应该是读取和修改、修改和修改之间，读和读是没有互斥操作的必要的。
//
//因此，衍生出另外一种锁，叫做读写锁。
//
//读写锁可以让多个读操作并发，同时读取，但是对于写操作是完全互斥的

//全局变量
var count int
var rLock sync.RWMutex

func Read(i int) {
	rLock.RLock()
	fmt.Printf("读 goroutine%d 数据=%d\n", i, count)
	defer rLock.RUnlock()
}

func Write(i int) {
	rLock.Lock()
	count = rand.Intn(1000)
	fmt.Printf("写 goroutine%d 数据=%d\n", i, count)
	defer rLock.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go Write(i)
	}
	for i := 0; i < 5; i++ {
		go Read(i)
	}
	time.Sleep(time.Second * 2)
}
