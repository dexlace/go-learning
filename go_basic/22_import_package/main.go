package main

// 包名的导入
// 包名是从$GOPATH/src/后开始计算的  使用/进行路径分割

// 导入包语句会自动触发包内部init（）函数的调用
// init函数没有参数也没有返回值

// 被最后导入的包会最先初始化并调用其init（）函数

import (
	"fmt"
	//"go_basic/21_package/calc"
	// 起别名
	doglast "go_basic/21_package/calc"
)

func main() {
	//fmt.Printf("%v", calc.Add(2, 7))
	fmt.Printf("%v\n", doglast.Add(2, 7))
	//fmt.Printf("%v", calc.Name)
	fmt.Printf("%v\n", doglast.Name)

}
