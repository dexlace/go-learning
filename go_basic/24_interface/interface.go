package main

import "fmt"

//type 接口名 interface {
//	方法名1 (参数列表1) 返回值列表1
//	方法名2 (参数列表2) 返回值列表2
//}

// 使用type将接口定义为自定义的类型名
// Go语言的接口在命名时，一般会在单词后面添加er
// 如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等，接口名最好突出该接口的类型含义

// 当方法名首字母是大写且这个接口类型名首字母也是大写时，
//这个方法可以被接口所在的包（package）之外的代码访问

//参数列表和返回值列表中的参数变量名是可以省略
// 接口是一种类型

type Speaker interface {
	speak()
}

type Dog struct {
}

// dog结构体的方法
func (d Dog) speak() {
	fmt.Println("汪汪汪")
}

type Cat struct {
}

// cat结构体的方法
func (c Cat) speak() {
	fmt.Println("喵喵喵")
}

// 以上两个结构体都实现了speak()方法，那么就可以作为结构体Speaker的一种类型

// 下面定义一个函数
func da(s Speaker) {
	s.speak()
}

// 2. 接口的实现
// 一个类型如果实现了接口中规定的所有方法，那么这个类型可以被理解为该接口类型
// 注意 是所有方法 可以将该类型赋值给接口  相当于java和C++的多态
// 注意是所有方法

// 3. 空接口
// 没有定义任何方法的接口就是空接口
// 空接口在实际项目中用的是非常多的，用空接口可以表示任意数据类型。
// 空接口也可以直接当做类型来使用，可以表示任意类型。相当于Java中的Object类型

type EmptyA interface {
}

// 空接口可以作为函数的参数，使用空接口可以接收任意类型的函数参数

func show(a interface{}) {
	fmt.Println(a)

}

func Print2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("其它类型")
	}
}

func main() {
	da(new(Dog))
	da(new(Cat))

	var dog Speaker = Dog{}
	dog.speak()

	var a EmptyA
	str := "hello"
	a = str
	fmt.Println(a)

	// 4.map的值为空接口
	var studentInfo = make(map[string]interface{})
	studentInfo["userName"] = "张三"
	studentInfo["age"] = 15
	studentInfo["isWork"] = true

	studentInfo["hobby"] = []string{"吃饭", "睡觉"}
	// 如何取得hobby对应的值
	// 用类型断言
	hobbies, ok := studentInfo["hobby"].([]string)
	if ok {
		fmt.Println(hobbies[0])
	}

	// 5 slice的值为空接口
	// 定义一个空接口类型的切片
	var slice = make([]interface{}, 4, 4)
	slice[0] = "张三"
	slice[1] = 1
	slice[2] = true

	// 6. 类型断言
	// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的
	// 如果我们想要判断空接口中值的类型，那么这个时候就可以使用类型断言，其语法格式：
	// x.(T)
	// X：表示类型为interface{}的变量
	// 表示断言x可能是的类型
	// 注意是接口值
	// 注意是接口值
	// 注意是接口值
	// 注意是接口值

	var ss interface{}
	ss = "12323"
	value, isString := ss.(string)
	if isString {
		fmt.Println("是String类型, 值为：", value)
	} else {
		fmt.Println("断言失败")
	}

}
