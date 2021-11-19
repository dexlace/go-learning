package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)







func getInput() string {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}


func main() {
	// 1. 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:23444")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	// 3.注意关闭连接
	defer conn.Close()
	// 2.写数据
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		//data, err := proto.Encode(msg)
		//if err != nil {
		//	fmt.Println("encode msg failed, err:", err)
		//	return
		//}
		// 2.1写数据
		conn.Write([]byte(msg))
	}
}