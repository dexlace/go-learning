package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(responseWriter http.ResponseWriter, request *http.Request) {
	// 1.解析url传递的参数，对于POST则解析响应包的主体（request body）
	request.ParseForm()
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(request.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	// 2. 拿到表单的数据
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 3. 这个写入到w的是输出到客户端的
	fmt.Fprintf(responseWriter, "Hello!erererer")
}

func login(response http.ResponseWriter, request *http.Request) {
	// 1. 拿到请求的方法
	fmt.Println("method:", request.Method) //获取请求的方法
	// 2. 请求的是显示登陆界面
	if request.Method == "GET" {
		// 5.0  防止表单重复提交
		// 在模版里面增加了一个隐藏字段token，这个值我们通过MD5(时间戳)来获取唯一值，
		// 然后我们把这个值存储到服务器端(session来控制
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))

		t, errors := template.ParseFiles("E:\\go_pro\\src\\go_web\\01_form_submit\\login.gtpl")
		fmt.Println(errors)
		// 无动态数据时
		//log.Println(t.Execute(response, nil))
		// 有动态数据时
		log.Println(t.Execute(response, token))
	} else {
		// 3. 请求的是登录数据，那么执行登录的逻辑判断
		request.ParseForm()
		//fmt.Println("username:", request.Form["username"])
		//fmt.Println("password:", request.Form["password"])
		token := request.Form.Get("token")
		// 5.1 防止表单重复提交
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}

		// 4 预防“跨站脚本攻击”（Cross Site Scripting, 安全专家们通常将其缩写成 XSS）“攻击
		//Go的html/template里面带有下面几个函数可以帮你转义
		//
		//func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
		//func HTMLEscapeString(s string) string //转义s之后返回结果字符串
		//func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
		fmt.Println("username:", template.HTMLEscapeString(request.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(request.Form.Get("password")))
		template.HTMLEscape(response, []byte(request.Form.Get("username"))) //输出到客户端

	}
}

func main() {

	//fmt.Println(GetCurrentDirectory())
	http.HandleFunc("/login", login)   //设置访问的路由
	http.HandleFunc("/", sayHelloName) //设置访问的路由

	err := http.ListenAndServe("127.0.0.1:9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
