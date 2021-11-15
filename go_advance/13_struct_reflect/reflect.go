package main

import (
	"fmt"
	"reflect"
)


// 任意值通过reflect.Typeof）获得反射对象信息后，如果它的类型是结构体，
//可以通过反射值对象（reflect.Type）的NumField（）
//和Field（）方法获得结构体成员的详细信息。

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (s Student) GetInfo() string{
	fmt.Println("这是get方法")
	fmt.Println(s.Name,s.Age)
	str:=fmt.Sprint("姓名：%v,年龄:%v",s.Name,s.Age)
	return str
}

func (s *Student)SetInfo(name string,age int)  {
	s.Name=name
	s.Age=age
}

func (s Student)PrintStudent()  {
	fmt.Println("这是print方法")
	fmt.Println(s.Name,s.Age)
	fmt.Println("打印学生")
}

// 一、获取结构体的属性

func getStructRunField(x interface{}){

	t := reflect.TypeOf(x)
	fmt.Println("类型",t)

	// 1. 判断传递过来的是否是结构体
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("请传入结构体类型!")
		return
	}

	//
	//| 方法                                        | 说明                                         |
	//| ------------------------------------------- | -------------------------------------------- |
	//| Field(i int)StructField                     | 根据索引，返回索引对应的结构体字段的信息     |
	//| NumField() int                              | 返回结构体成员字段数量                       |
	//| FieldByName(name string)(StructField, bool) | 根据给定字符串返回字符串赌赢的结构体字段信息 |
	//| FieldByIndex(index []int)StructField        | 多层成员访问时，根据[] int 提供的每个结构    |

	// 2. 通过类型变量里面的Field可以获取结构体的字段
	field0:=t.Field(0) // 获取第0个字段
	fmt.Printf("%#v \n", field0)
	fmt.Println("字段0名称:", field0.Name)
	fmt.Println("字段0类型:", field0.Type)
	fmt.Println("字段0Tag:", field0.Tag.Get("json"))



	// 3. 尝试获取结构体中的Age字段
	field1, ok := t.FieldByName("Age")
	if ok  {
		fmt.Println("字段1名称:", field1.Name)
		fmt.Println("字段1类型:", field1.Type)
		fmt.Println("字段1Tag:", field1.Tag)
	}

	// 4. 通过类型变量里面的NumField获取该结构体有几个字段
	var fieldCount = t.NumField()
	fmt.Println("结构体有：", fieldCount, " 个属性")

	// 5. 获取值
	v:=reflect.ValueOf(x)
	nameValue:=v.FieldByName("Name")
	fmt.Println("nameValue:",nameValue)
	ageValue:=v.FieldByIndex([]int{1})
	fmt.Println("ageValue:",ageValue)


}

// 二、获取方法 这里打算传入的是结构体值类型

func getStructFn(s interface{}){
	t := reflect.TypeOf(s)
	fmt.Println("----------判断是否是结构体------")
	// 判断传递过来的是否是结构体
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("请传入结构体类型!")
		return
	}

	// 1. 通过类型变量里面的Method，可以获取结构体的方法
	fmt.Println("---------获取第一个方法名-------")
	method0 := t.Method(0)
	// 获取第一个方法， 这个是和ACSII相关
	fmt.Println(method0.Name)

	// 2. 通过类型变量获取这个结构体有多少方法
	// 明明是由三个方法  但是打印出来只有两  因为set方法是引用传递 无法调用
	fmt.Println("---------获取方法数-------")
	methodCount := t.NumMethod()
	fmt.Println("拥有的方法", methodCount)

	// 3. 通过值变量 执行方法（注意需要使用值变量，并且要注意参数）
	// 自动传值
	fmt.Println("---------自动传值调用方法一-------")
	v := reflect.ValueOf(s)
	// 拿到PrintStudent并执行
	v.MethodByName("PrintStudent").Call(nil)
	v.MethodByName("GetInfo").Call(nil)
	fmt.Println("---------自动传值调用方法二-------")
	v.Method(0).Call(nil)
	v.Method(1).Call(nil)

}


// 三、获取方法 传入的是结构体引用类型
// 形参不变 但这里打算传入结构体的引用类型

func getStructSetFn(s interface{}) {
	//获取s的指针的reflect.Type
	ptrType := reflect.TypeOf(s)
	trueType:=ptrType.Elem()
	fmt.Println("----------判断是否是结构体------")
	// 1. 判断传递过来的是否是结构体
	if ptrType.Kind() != reflect.Struct && trueType.Kind() != reflect.Struct {
		fmt.Println("请传入结构体类型!")
		return
	}

	// 1. 通过类型变量里面的Method，可以获取结构体的方法
	fmt.Println("---------获取第一个方法名-------")
	method0 := ptrType.Method(0)
	// 获取第一个方法， 这个是和ACSII相关
	fmt.Println(method0.Name)

	// 2. 通过类型变量获取这个结构体有多少方法
	// 这里可以打印出有三个方法  因为传入的是引用  能运行set方法
	fmt.Println("---------获取方法数-------")
	methodCount := ptrType.NumMethod()
	fmt.Println("拥有的方法", methodCount)
	//
	// 3. 通过值变量 执行方法（注意需要使用值变量，并且要注意参数）
	// 自动传值
	fmt.Println("---------自动传值调用方法-------")
	v := reflect.ValueOf(s)
	// 拿到PrintStudent并执行
	v.MethodByName("PrintStudent").Call(nil)
	v.MethodByName("GetInfo").Call(nil)

	fmt.Println("---------手动传参调用方法-------")
	// 4. 手动传参
	var params []reflect.Value
	params = append(params, reflect.ValueOf("张三"))
	params = append(params, reflect.ValueOf(23))
	// 执行setInfo方法
	v.MethodByName("SetInfo").Call(params)

	// 通过值变量来获取参数
	v.MethodByName("PrintStudent").Call(nil)



}



func main() {

	st:=Student{
		"zhangsan",
		21,
	}
	getStructRunField(st)

	fmt.Println("................值传递：获取结构体的方法.........")
	getStructFn(st)

	fmt.Println("................引用传递：获取结构体的方法二.........")
	st2:=&st
	getStructSetFn(st2)

}
