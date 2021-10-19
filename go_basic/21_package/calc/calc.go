package calc

// 自定义包，最好和文件夹统一起来

// 私有 不可被引用

var age = 10

// 公有变量 可以被外部引用到

var Name = "张三"

// 首字母大写，表示共有方法

func Add(x, y int) int {
	return x + y
}
func Sub(x, y int) int {
	return x - y
}
