package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("\n/********************第10行 map初始化****************/\n")
	// make常用来分配内存
	var userInfo = make(map[string]string)
	userInfo["name"] = "zhangsan"
	userInfo["age"] = "20"
	userInfo["sex"] = "male"
	fmt.Println(userInfo)
	fmt.Println(userInfo["name"])

	var userInfo2 = map[string]string{
		"username": "lisi",
		"age":      "23",
		"sex":      "female",
	}
	fmt.Println(userInfo2)

	fmt.Printf("\n/********************第26行 map的遍历****************/\n")
	for key, value := range userInfo2 {
		fmt.Println("key:", key, " value:", value)
	}

	fmt.Printf("\n/********************第31行 判断map某个键值是否存在****************/\n")
	// 我们在获取map的时候，会返回两个值，也可以是返回的结果，一个是是否有该元素
	value, ok := userInfo2["age2"]
	fmt.Println(value, ok)

	fmt.Printf("\n/********************第36行 删除键值对****************/\n")
	delete(userInfo2, "age")
	fmt.Println(userInfo2)

	fmt.Printf("\n/********************第40行 元素为map类型的切片****************/\n")
	// 切片在中存放map
	var userInfoList = make([]map[string]string, 3, 3)
	var user = map[string]string{
		"userName": "张安",
		"age":      "15",
	}
	var user2 = map[string]string{
		"userName": "张2",
		"age":      "15",
	}
	var user3 = map[string]string{
		"userName": "张3",
		"age":      "15",
	}
	userInfoList[0] = user
	userInfoList[1] = user2
	userInfoList[2] = user3
	fmt.Println(userInfoList)

	for _, item := range userInfoList {
		fmt.Println(item)
	}

	fmt.Printf("\n/********************第64行 值为切片类型的map****************/\n")

	userInfo3 := make(map[string][]string)
	userInfo3["hobby"] = []string{"吃饭", "睡觉", "上班"}
	fmt.Println(userInfo3)

	st1 := "how are u"
	arr := strings.Split(st1, " ")
	fmt.Println(arr)
	fmt.Println(arr[0])

	countMap := make(map[string]int)
	for _, v := range arr {
		countMap[v]++
	}
	fmt.Println(countMap)

}
