package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	// 切片是引用数据类型
	// 所以牵一发动全身
	var enum = []int{100, 200, 300}
	fnum := enum
	fnum[0] = 800
	fmt.Println(enum, fnum)

	// 声明切片的方式与数组的方式除了容量没有别的
	// 因为他支持自动扩容  所以没有容量要求  且他是一个引用类型
	// var name [] T
	var slice []int
	fmt.Println(slice == nil)

	var slice2 = []int{1, 2, 3}
	fmt.Println(slice2)

	for i := 0; i < len(slice2); i++ {
		fmt.Println(slice2[i], " ")
	}

	fmt.Printf("\n/********************第26行 基于数组定义切片****************/\n")

	// 基于数组定义切片 sa是一个数组
	arraya := [5]int{10, 20, 30, 40, 50}
	// 获取数组所有值,返回的是一个切片
	slicea := arraya[:]
	fmt.Println(slicea)

	// 获取[1,4)的数据
	sliceb := arraya[1:4]
	fmt.Println(sliceb)

	// 获取[0,3)
	slicec := arraya[:3]
	fmt.Println(slicec)

	// 获取[0,3)
	sliced := arraya[3:]
	fmt.Println(sliced)

	fmt.Printf("\n/********************第46行 切片的长度和容量****************/\n")

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d \n", len(s), cap(s))
	sss := s[2:]
	fmt.Printf("长度%d 容量%d \n", len(sss), cap(sss))

	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	// 经过切片后sss = [5, 7] 所以切片的长度为2，但是一因为容量是从2的位置一直到末尾，所以为4
	ssss := s[2:4]
	fmt.Printf("长度%d 容量%d \n", len(ssss), cap(ssss))

	var slll = make([]int, 4, 8)
	fmt.Printf("长度%d 容量%d \n", len(slll), cap(slll))

	fmt.Printf("\n/********************第61行 切片扩容****************/\n")
	// 给切片扩容
	sli2 := []int{1, 2, 3, 4}
	sli2 = append(sli2, 5)
	fmt.Println(sli2)

	fmt.Printf("\n/********************第67行 切片合并****************/\n")
	// 给切片合并
	sli3 := []int{6, 7, 8}
	sli2 = append(sli2, sli3...)
	fmt.Println(sli2)

	fmt.Printf("\n/********************第73行 切片复制****************/\n")
	// 切片复制 用copy可以不改变源切片值
	sli4 := []int{1, 2, 3, 4}
	// 使用make创建一个切片
	sli5 := make([]int, len(sli4), len(sli4))
	fmt.Println(sli5)
	copy(sli5, sli4)
	sli5[3] = 5
	fmt.Println(sli4)
	fmt.Println(sli5)

	fmt.Printf("\n/********************第84行 切片中的值删除****************/\n")
	sli6 := []int{0, 1, 2, 3, 4, 5}
	sli6 = append(sli6[:1], sli6[3:]...)
	fmt.Println(sli6)

	fmt.Printf("\n/********************第89行 切片比较****************/\n")
	aaaaa := []byte{0, 1, 3, 2}

	bbbbb := []byte{0, 1, 3, 2}

	ccccc := []byte{1, 1, 3, 2}

	// 只适用于byte
	fmt.Println(bytes.Equal(aaaaa, bbbbb))

	fmt.Println(bytes.Equal(aaaaa, ccccc))

	// 用反射比较会比较通用
	fmt.Println(reflect.DeepEqual(aaaaa, bbbbb))

	fmt.Println(reflect.DeepEqual(aaaaa, ccccc))

}
