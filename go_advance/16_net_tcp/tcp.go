package main

import (
	"fmt"
	"net"
)

// 用于接收请求的方法
func processConn(conn net.Conn)  {
	// 与客户端通信
	var tmp [128]byte
	// 使用for循环监听消息
	for {
		// 4. 即读数据
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(conn, string(tmp[:n]))
	}
}
func main() {

	// 1. tcp服务端本地端口启动服务
	listen,err:=net.Listen(
		"tcp",
		"127.0.0.1:23444",
	)

	if err!=nil{
		fmt.Println("start server failed",err)
	}


	// 2. for循环监听
	for {
		// 2. 等待别人来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			return
		}
		// 3. 与客户端通信
		go processConn(conn)
	}


}
