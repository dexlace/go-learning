package main

import (
	"fmt"
	"sync"
	"time"
)

// 能保证once只执行一次，无论你是否更换once.Do(xx)这里的方法,
//这个sync.Once块只会执行一次。
var once sync.Once

func main() {

	for i := 0; i < 10; i++{
		// once里面传的是函数指针
		once.Do(onceYes)
		// onceYes只执行一次
		fmt.Println("count:", "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			// 由于once.Do已经执行过onceYes的方法了，所以这里的onceNo是
			// 无法执行的
			// 当然，不影响以下的代码执行
			once.Do(onceNo)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onceYes() {
	fmt.Println("onceYes")
}
func onceNo() {
	fmt.Println("onceNo")
}
