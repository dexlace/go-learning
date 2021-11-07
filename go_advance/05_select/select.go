package main

import "fmt"

//  在某些场景下我们需要同时从多个通道接收数据
// select的使用类似于switch 语句
// 它有一系列case分支和一个默认的分支。
// 每个case会对应一个管道的通信（接收或发送）过程。
// select会一直等待，直到某个case的通信操作完成时，
// 就会执行case分支对应的语句

func main() {
	cha1 := make(chan int, 10)
	cha1 <- 10
	cha1 <- 11
	cha1 <- 12
	cha1 <- 13
	cha2 := make(chan int, 10)
	cha2 <- 21
	cha2 <- 22
	cha2 <- 23
	cha2 <- 24
	for {
		select {
		case value := <-cha1:
			fmt.Println("从cha1中读取数据：", value)
		case value := <-cha2:
			fmt.Println("从cha2中读取数据：", value)
		default:
			fmt.Println("所有数据获取完毕")
			return
		}

	}

}
