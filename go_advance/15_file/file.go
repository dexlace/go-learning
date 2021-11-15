package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func main() {

	// 1.打开文件
	// - name：要打开的文件名
	//- flag：打开文件的模式
	//  - os.O_WRONLY：只读
	//  - os.O_CREATE：创建
	//  - os.O_RDONLY：只读
	//  - os.O_RDWR：读写
	//  - os.O_TRUNC：清空
	//  - os.O_APPEND：追加
	//- perm：文件权限，一个八进制数，r（读）04，w（写）02，x（执行）01
	file, _ := os.OpenFile("E:\\test2.txt", os.O_CREATE | os.O_RDWR | os.O_APPEND, 777)
	defer file.Close()
	//str := "啦啦啦 \r\n"
	// 2. 写入文件
	//file.WriteString(str)

	// 3.第二方式通过bufio写入
	writer := bufio.NewWriter(file)
	// 先将数据写入缓存
	writer.WriteString("你好，我是通过writer写入的 \r\n")
	// 将缓存中的内容写入文件
	writer.Flush()

	// 第三种方式，通过ioutil
	str2 := "hello"
	ioutil.WriteFile("./main/test.txt", []byte(str2), 777)
}
