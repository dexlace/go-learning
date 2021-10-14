package main

import "fmt"

func main(){

	fmt.Printf("\n/********************第7行 重点重点重点之数组初始化****************/\n")

	// 数组的长度是类型的一部分
	var arr1 [3]int
	var arr2 [4]int
	fmt.Printf("%T,%T\n", arr1, arr2)

	// 数组的初始化方法
	var arr3 [3]int
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3
	fmt.Println(arr3)

	// 第二种初始化数组的方法
	// 数组的打印还是比较银杏的
	var arr4 = [4]int{10, 20, 30, 40}
	fmt.Println(arr4)

	// 第三种数组初始化方法，自动推断数组长度
	var arr5 = [...]int{12, 22}
	fmt.Println(arr5)

	// 第四种初始化数组的方法  指定下标
	arr6 := [...]int{1: 1, 3: 5}
	fmt.Println(arr6)

	fmt.Printf("\n/********************第34行 重点重点重点之遍历数组****************/\n")
	anum := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(anum); i++ {
		fmt.Println(anum[i], " ")
	}

	// _通常用于不需要使用但必须。。。的。。。
	// 比如下面的数组遍历
	bnum := [...]int{10, 20, 30, 40}
	for _, v := range bnum {
		fmt.Print(v, " ")
	}

	fmt.Printf("\n/********************第47行 重点重点重点之数组的元素属于值类型****************/\n")
	// java中是属于引用数据类型
	var cnum = [...]int{111, 222, 333}
	dnum := cnum
	dnum[0] = 444
	fmt.Println(cnum, dnum)

	fmt.Printf("\n/********************第54行 重点重点重点之二维数组****************/\n")
	var tnum = [2][2]int{{1, 1}, {2, 2}}
	fmt.Println(tnum)

	var ttnum = [2][2]int{{1, 1}, {3, 3}}
	for i := 0; i < len(ttnum); i++ {
		for j := 0; j < len(ttnum[0]); j++ {
			fmt.Println(ttnum[i][j])
		}
	}

	// for range遍历
	for _, item := range ttnum {
		for _, item2 := range item {
			fmt.Println(item2)
		}
	}

}
