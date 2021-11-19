package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 1. 代替浏览器去发请求  是一个客户端 浏览器就是客户端
	res, err := http.Get("http://127.0.0.1:9090/query?name=zansan&age=10")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2.从res中把服务端返回的数据读取出来
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// 2. 请求方法二
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/query")
	data.Set("name", "周林")
	data.Set("age", "100")
	// 对请求进行编码
	queryStr := data.Encode()
	urlObj.RawQuery = queryStr

	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req)

	// todo 待续。。。
}
