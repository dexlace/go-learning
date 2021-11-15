package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 1. 用file的read读取文件
func fileRead(){
	// 1. 打开文件
	file, err := os.Open("E:\\test.txt")
	// 2. 关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件出错")
	}

	// 3. 读取文件里面的内容
	var tempSlice = make([]byte, 1024)
	var strSlice []byte
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Printf("读取完毕")
			break
		}
		fmt.Printf("读取到了%v 个字节 \n", n)
		strSlice := append(strSlice, tempSlice...)
		fmt.Println(string(strSlice))
	}
}

// 2. 通过bufio方式读取

func bufioRead(){
	// 1. 打开文件
	file, err := os.Open("E:\\test.txt")
	// 2. 关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件出错")
	}

	// 3. 是个指针对象
	fmt.Printf("file=%v\n",file)

	// 4. 创建一个reader 带有缓冲区 默认大小为4096
	reader:=bufio.NewReader(file)
	// 用来接收文件内容
	var fileStr string
	var count int = 0
	for {
		// 5.ReadString
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			// 读取完成的时候，也会有内容
			fileStr += str
			fmt.Println("读取结束", count)
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		// 读取
		count ++
		fileStr += str
	}

	fmt.Println(fileStr)

}

// 3.文件读取方式三 文件比较少的时候，可以通过ioutil来读取文件
func ioutilRead()  {
	// 通过IOUtil读取
	byteStr, _ := ioutil.ReadFile("E:\\test.txt")
	fmt.Println(string(byteStr))

}

func main() {

	//
	//fileRead()

	ioutilRead()



}
