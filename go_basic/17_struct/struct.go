package main

import (
	"fmt"
)

// 结构体
// 结构体首字母可以大写也可以小写，
//大写表示这个结构体是公有的，在其它的包里面也可以使用，
//小写表示结构体属于私有的，在其它地方不能使用
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 实例化方式一
	var p person
	p.name = "zhangsan"
	p.age = 19
	p.gender = "male"
	p.hobby = []string{"篮球", "足球", "罗翔"}
	fmt.Println(p)
	// 访问p的类型
	fmt.Printf("%#v\n", p)
	fmt.Println(p.name)

	fmt.Println("-----------实例化方式二--------")
	// 实例化方式二  用指针
	var p2 = new(person)
	p2.name = "lisi"
	// 等价于(*p2).name="lisi"
	p2.age = 12
	fmt.Printf("%#v\n", p2)

	fmt.Println("-----------实例化方式三--------")
	// 实例化方式三 用&取地址
	var p3 = &person{}
	p3.name = "wangwu"
	fmt.Printf("%#v\n", p3)

	fmt.Println("-----------实例化方式四--------")
	// 键值对的方式来实例化结构体，实例化的时候，可以直接指定对应的值
	// 注意赋值时的逗号  定义结构体时没有逗号
	// 实例化方式四
	var p4 = person{
		name: "zhaoliu",
		age:  10,
	}

	fmt.Printf("%#v\n", p4)

	fmt.Println("-----------实例化方式五--------")
	var p5 = &person{
		name: "孙五",
		age:  10,
	}
	fmt.Printf("%#v\n", p5)

	fmt.Println("-----------实例化方式六--------")
	// 这里要完全赋值
	var p6 = person{
		 "陈其",
		  10,
		  "male",
		[]string{"chifan","shuijiao","shangban"},
	}
	fmt.Printf("%#v\n", p6)

	fmt.Println("-----------匿名结构体--------")
	var s struct {
		age  int
		name string
	}
	s.name = "sds"

}
