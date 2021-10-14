package main

import "fmt"

// Go语言中函数
func f(x int, y int) int {
	return x + y
}

// 没有返回值的函数
func f1(x int, y int) {
	fmt.Println("sdsd")
}

// 没有参数
func f2() {
	fmt.Println("drerer")
}

// 没有参数但有返回值
func f3() int {
	return 3
}

//函数定义时可以给返回值命名
func sum(x int, y int) (ret int) {
	ret = x + y
	// 有返回值时可以不写return 返回值 直接写return就好
	//return ret
	return

}

//支持多返回值
func f12(x int, y int) (sum int, dec int) {
	sum = x + y
	dec = x - y
	return
}

func f13() (int, string) {
	return 1, "nihao"
}

// 参数的类型简写 x,y都是同样的int类型
func f14(x,y int) int{
	return x+y
}

func f15(x string,m,n int) int{
	return m+n
}

// 可变长参数
// 可变长参数必须放在最后
func f16(x string,y ...int){
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f13()
	fmt.Println(n)

	sum,dec:=f12(15,16)
	fmt.Println(sum,dec)

	f16("nihao",1,2,3,4,5)
	f16("nih")






}
