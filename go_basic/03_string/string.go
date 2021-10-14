package main

import (
	"fmt"
	"strconv"
)

func main()  {

	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("一个汉字长度竟然为3,你好的长度为%d\n", len(s2))
	// 字符串连接
	fmt.Println(s1 + s2)
	// 字符串长度
	fmt.Println(len((s1 + s2)))

	// 多行字符串  用反引号
	var str = `第一行
第二行`
	fmt.Println(str)

	/**
	  - strings.Split：分割
	  - strings.contains：判断是否包含
	  - strings.HasPrefix，strings.HasSuffix：前缀/后缀判断
	  - strings.Index()，strings.LastIndex()：子串出现的位置
	  - strings.Join()：join操作
	  - strings.Index()：判断在字符串中的位置
	*/


	var ba byte = 'a'
	// %v是原样输出  记住
	fmt.Printf("原样输出字符会显示该字符的ASCII码值 %v\n", ba)
	// %c是才输出字符 记住
	fmt.Printf("原样输出字符会显示该字符的ASCII码值 %c\n", ba)

	// for循环打印字符串里的字符
	//ss :="abcdefg"
	//ss :="abc defg"
	ss := "abc你fg"
	for i := 0; i < len(ss); i++ {
		fmt.Printf("%v(%c)\t", ss[i], ss[i])
	}

	// 这里表示的是一个UTF-8字符
	// 当需要处理中文，日文或者其他复合字符时，则需要用到rune类型，rune类型实际上是一个int32
	for index, value := range ss {
		//fmt.Println(index,v)
		// 以下才能愉快的打出“你”
		fmt.Printf("%v  %c\n", index, value)
	}

	fmt.Println("/*******************第54行 字符串转换成字符数组*************************/")
	ss2 := "zifuchuan"
	// 不能直接修改字符串 需要先转换为字符数组后再修改
	byteS1 := []byte(ss2)
	// 修改
	byteS1[0] = 'h'
	// 重新生成一个
	fmt.Println(string(byteS1))

	ss3 := "你好啊gooo"
	byteS2 := []rune(ss3)
	byteS2[2] = '吗'
	fmt.Println(string(byteS2))


	fmt.Println("/*************************第69行 fmt.Sprintf()函数实现到字符串类型的转换**********************************/")
	var ii int = 20
	var ff float64 = 12.3456
	var tt bool = true
	var bbb byte = 'a'
	strii := fmt.Sprintf("%d", ii)
	// %T表示输出数据的类型
	fmt.Printf("类型：%v-%T \n", strii, strii)

	strff := fmt.Sprintf("%f", ff)
	// %T表示输出数据的类型
	fmt.Printf("类型：%v-%T \n", strff, strff)

	strtt := fmt.Sprintf("%t", tt)
	// %T表示输出数据的类型
	fmt.Printf("类型：%v-%T \n", strtt, strtt)

	strbb := fmt.Sprintf("%c", bbb)
	// %T表示输出数据的类型
	fmt.Printf("类型：%v-%T \n", strbb, strbb)

	fmt.Println("/*************************第93行 strconv实现到字符串类型的转换**********************************/")
	var nnn1 int64 = 20
	// 第二个参数表示将该数以多少进制解析
	snum1 := strconv.FormatInt(nnn1, 10)
	fmt.Printf("转换：%v - %T\n", snum1, snum1)

	var nnn2 float64 = 4.12344
	/*
		参数1：要转换的值
		参数2：格式化类型 'f'表示float，'b'表示二进制，‘e’表示 十进制
		参数3：表示保留的小数点，-1表示不对小数点格式化
		参数4：格式化的类型，传入64位 或者 32位
	*/
	snum2 := strconv.FormatFloat(nnn2, 'f', -1, 64)
	fmt.Printf("转换：%v - %T\n", snum2, snum2)


	fmt.Println("/***第111行 字符串转换成int和float类型  暂时不管  暂时没理解这里的接口概念***/")
	//strrrr :="10000"
	// 第一个参数：需要转换的数  第二个参数：进制  第三个参数  32位或者64位

	//int64, err := strconv.ParseInt(strrrr, 10, 64)
}
