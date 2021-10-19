package main

import (
	"fmt"
	"time"
)

func main() {

	a := time.Now()
	fmt.Println(a)

	year := a.Year()
	month := a.Month()
	day := a.Day()
	fmt.Printf("%d-%02d-%02d \n", year, month, day)

	// 格式化时间
	// 注意go语言的Format不是从1970年算的
	// 而是go语言的诞生时间2006年1月2日 15点04分
	// 24小时制
	fmt.Println(a.Format("2006-01-02 15:04:05"))
	// 12小时制
	fmt.Println(a.Format("2006-01-02 03:04:05"))

	// 其他略 没啥必要 工具性的东西

}
