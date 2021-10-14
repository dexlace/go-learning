package main

import "fmt"

// defer 把他后面的语句延迟到函数即将返回时再执行
// 一个函数中可以有多个defer 按照先进后出的顺序执行
func deferDemo()  {
 fmt.Println("start")
 defer fmt.Println("关闭资源")
 defer fmt.Println("关闭资源2")
 fmt.Println("end" )

}

// go语言中函数的return不是原子操作，在底层分为两步来执行
// 第一步： 返回值赋值
// 第二步： 真正的ret返回
// 如果存在defer，那么defer执行的时机在第一步和第二步之间
func f1()int{
	x:=5
	defer func(){
		// 修改的是x 不是返回值
		x++
	}()
	// x=5
	// 然后执行return
	//   return操作分为两步，首先返回值赋值为5，这里的返回值没有显示声明
	//                        所以赋值之后的返回值永远是5
	// 即使defer语句将x修改   但是修改的是x，不是返回值
	// x与返回值之间仅仅是值传递
	return x
}


func f2()(x int){

	defer func(){
		x++
	}()
	// 首先执行返回值x=5的操作，注意，因为显示声明了x就是返回值
	// 所以defer语句的插入影响到了x的值，x的值即返回值
	// 所以返回的值为6
	return 5
}

// 按照f1的思路就很好理解
func f3()(y int)  {
	x:=5
	defer func() {
		x++
	}()
	// y=x=5
	return x
}

func f4()(x int){
	// 参数x被f4的返回值x传参，是值传递
	// 两个x只是值传递
	defer func(x int) {
		x++
	}(x) // 将x（返回值）赋值给匿名函数的参数 返回值不变 还是5
	return 5
}


func main() {
	deferDemo()
	fmt.Println("---------defer的执行时机---------")
	x:=f1()
	fmt.Println("f1的结果是",x)

	fmt.Println("f2的结果是",f2())

	fmt.Println("f3的结果是",f3())

	fmt.Println("f4的结果是",f4())


}