package main

import (
	"encoding/json"
	"fmt"
)

// 1. 序列化与反序列化
// Golang JSON序列化是指把结构体数据转化成JSON格式的字符串
// Golang JSON的反序列化是指把JSON数据转化成Golang中的结构体对象
// Golang中的序列化和反序列化主要通过“encoding/json”包中的 json.Marshal()
// 和 son.Unmarshal()

// 注意需要序列化的
// 结构体字段的首字母需要大写
// 表示公有
// 否则无法转换

// 2. 结构体标签Tag
// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来
// Tag在结构体字段的后方定义，由一对反引号包裹起来
// 注意事项：为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，
// 编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

type student struct {
	Id     string `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	grade  string // 所以这个字段是无法序列化的
}

// 3. 嵌套结构体的序列化和反序列化和普通结构体类似 不再多说

func main() {
	s1 := student{
		"12",
		"male",
		"zhangsan",
		"class 1",
	}
	// 结构体转换为json 返回的是byte类型的切片
	bytes, _ := json.Marshal(s1)
	str := string(bytes)
	fmt.Println(str)

	// // Json字符串转换成结构体
	var str2 = `{"Id":"12","Gender":"男","Name":"李四","grade":"class 2"}`
	var s2 student
	// 第一个是需要传入byte类型的数据，第二参数需要传入转换的地址
	err := json.Unmarshal([]byte(str2), &s2)
	if err != nil {
		fmt.Printf("转换失败 \n")
	} else {
		fmt.Printf("%#v \n", s2)

	}
}
