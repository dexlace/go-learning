package main

import (
	"fmt"
	"reflect"
)

// 学过java 当然知道反射是什么
// 反射就是在程序运行期间对程序本身进行访问和修改的能力

func getRun(x interface{}){
	// 1.使用reflect.TypeOf（）函数可以接受任意interface{}参数，
	// 可以获得任意值的类型对象（reflect.Type）
	typeVal :=reflect.TypeOf(x)
	fmt.Println("运行期间的类型是", typeVal)

	// 2.reflect.ValueOf() 返回的是reflect.Value类型
	valueVal:=reflect.ValueOf(x)
	fmt.Println("运行期间的值是", valueVal)

	// 3.1 将reflect.Value类型转化成interface{},也就是空接口
	iV:=valueVal.Interface()
	//  通过断言可以转换成需要的类型
	t1,isSuccess:=iV.(int)
	fmt.Println("t1=",t1,"是否转换成功",isSuccess)

	// 3.2 获取reflect.Value的原始值
	//  注意也是要做断言下去获取原始值
	if isSuccess{
		fmt.Println("reflect.Value类型的原始类型为Int，其原始值为",valueVal.Int())
	}

}


// 4. 关于反射中的type name和type kind
//  底层的有name和kind
// 自定义的比如数组 切片 map 指针 他们的.Name()为空
func getRun2(x interface{}){
	v:=reflect.TypeOf(x)
	fmt.Println("类型",v)
	fmt.Println("类型名称",v.Name())
	fmt.Println("类型种类",v.Kind())

}


type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (s Student) GetInfo() string{
	str:=fmt.Sprint("姓名：%v,年龄:%v",s.Name,s.Age)
	return str
}

func (s *Student)SetInfo(name string,age int)  {
	s.Name=name
	s.Age=age
}

func getStructRun(x interface{}){


}

func main() {

	// reflect.TypeOf()获取任意值的类型对象
	getRun(10)
	getRun(10.01)
	getRun("abc")
	getRun(true)

	fmt.Println("/****************type name和type kind***********************/")
	getRun2(900)
	// 像数组、切片、Map、指针等类型的变量，它们的.Name（）都是返回空。
	var arr4 = [4]int{10, 20, 30, 40}
	getRun2(arr4)



}
