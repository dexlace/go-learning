package main

import "fmt"

func main()  {
	// 注意++和--操作只能单独使用
	// 而且注意只有i++ 没有++i
	var jjj int = 1
	jjj++

	fmt.Println("/********************if的流程控制****************/")

	// if后面推荐不使用括号
	var nuuu = 10
	if nuuu == 10 {
		fmt.Println("nuuu==10")
	} else if nuuu > 10 {
		fmt.Println("nuuu>10")
	} else {
		fmt.Println("nuuu <10 ")
	}

	// if 中用局部变量
	if numm := 11; numm >= 10 {
		fmt.Println("nummmmmmmmmmm")

	}

	fmt.Println("/********************for的流程控制****************/")

	// for的流程控制
	// for 初始语句; 条件表达式; 结束语句 {
	//	循环体
	//}
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i+1)
	}

	fmt.Printf("\n/********************go中没有while循环 但是可以通过for去实现****************/\n")
	/**
	for {
	    循环体
	}
	*/

	fmt.Printf("\n/********************for range的流程控制****************/\n")

	// 数组、切片、字符串返回索引和值。
	// 字符串遍历
	var strt = "你好golang"
	for key, value := range strt {
		fmt.Printf("%v-%c  ", key, value)
	}

	// 遍历切片（数组）
	var arrayy = []string{"php", "java", "c", "python"}
	for index, value := range arrayy {
		fmt.Printf("%v-%s ", index, value)
	}

	fmt.Printf("\n/********************switch case的流程控制****************/\n")

	// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
	casename := ".html"
	switch casename {
	case ".html":
		{
			fmt.Println(".html")
			break
		}
	case ".doc":
		{
			fmt.Println(".doc")
			break
		}
	case ".js":
		{
			fmt.Println(".js")
			break
		}
	default:
		{
			fmt.Println("其他后缀")

		}

	}

	// switch 的另外一种写法
	switch caseeee := ".txt"; caseeee {
	case ".html":
		{
			fmt.Println(".html")
			break
		}
	case ".doc":
		{
			fmt.Println(".doc")
			break
		}
	case ".js":
		{
			fmt.Println(".js")
			break
		}
	default:
		{
			fmt.Println("其他后缀")

		}

	}

	// switch的一个分支可以有很多值
	fmt.Printf("\n/********************switch的一个分支可以有很多值****************/\n")
	fmt.Printf("\n/********************switch的break可以不写****************/\n")

	ccccase := ".txt"
	switch ccccase {
	case ".html":
		{
			fmt.Println(".html")
			break
		}
	case ".txt", ".doc":
		{
			fmt.Println("传递来的是文档")
			break
		}
	case ".js":
		{
			fmt.Println(".js")
		}
	default:
		{
			fmt.Println("其它后缀")
		}

	}

	fmt.Printf("\n/********************第355行 switch的穿透: fallthrought****************/\n")

	// fallthrought 只能穿透紧挨着的一层，不会一直穿透，但是如果每一层都写的话，就会导致每一层都进行穿透
	cname := ".txt"
	switch cname {
	case ".html":
		fmt.Println(".html")
		fallthrough
	case ".txt", ".doc":
		fmt.Println("传递来的是文档")
		fallthrough
	case ".js":

		fmt.Println(".js")
		fallthrough

	default:
		fmt.Println("其它后缀")
	}

}
