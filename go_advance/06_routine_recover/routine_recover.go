package main

import "fmt"
// 如果不进行recover，便会导致整个程序挂掉，
//Recover()用法是：将Recover()写在defer中，
//并且在可能发生panic的地方之前，先调用此defer的东西
//（让系统方法域结束时，有代码要执行。）
//当程序遇到panic的时候（当然，也可以正常的调用出现的异常情况），
//系统将跳过后面的代码，进入defer，
//如果defer函数中recover()，则返回捕获到的panic的值。

func myrecover(x int){
	// 设置recover，recover只能放在defer之后
	defer func(){
		if thePanic:=recover(); thePanic!=nil{
			fmt.Println(thePanic)
		}
	}()

	var a[10] int
	a[x]=11
	fmt.Println("此函数到此是否执行，答案是不会")
}

func mytest(){
	fmt.Println("这里是否能执行，答案是能")
}



func main() {
	myrecover(11)
	mytest()

}
