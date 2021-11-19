package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http服务器


func f1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	// 识别URL中的参数
	fmt.Println(r.URL.Query())
	str := "hello world！"
	// 3. 用writer去写
	w.Write([]byte(str))
}

// 也就是这里演示的是request和response的用法

func f2(w http.ResponseWriter, r *http.Request) {
	// 4.拿到请求参数
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	age := queryParams.Get("age")
	fmt.Println("传递来的name:", name)
	fmt.Println("传递来的age:", age)
	// 5. 拿到请求方法
	fmt.Println(r.Method)
	// 6. 拿到请求体
	fmt.Println(ioutil.ReadAll(r.Body))
	str := "hello world！"
	w.Write([]byte(str))
}

func main() {
	// 2. 路径  处理器 后面的/不能去掉
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/home/", f1)
	http.HandleFunc("/query/", f2)

	// 1. 监听并启动服务 传一个ip地址
	http.ListenAndServe("127.0.0.1:9090", nil)
}
