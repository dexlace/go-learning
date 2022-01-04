# go的包管理

## 1. Go包分类

Golang中的包可以分为三种：==1、系统内置包   2、自定义包   3、第三方包==

- **系统内置包**：Golang 语言给我们提供的内置包，引入后可以直接使用，如fmt、strconv、strings、sort、errors、time、encoding/json、os、io等。
- **自定义包**：开发者自己写的包
- **第三方包**：属于自定义包的一种，需要下载安装到本地后才可以使用。

## 2. Go管理工具go mod

在==Golang1.11版本之前如果我们要自定义包的话必须把项目放在GOPATH目录==。Go1.11版本之后无需手动配置环境变量，使用

==go mod 管理项目==，也不需要非得把项目放到GOPATH指定目录下，你可以在你磁盘的任何位置新建一个项目，Go1.13以后

可以彻底不要GOPATH了。

### 2.1 go mod init初始化项目

实际项目开发中我们首先要在我们项目目录中用==`go mod init`命令生成一个go.mod文件管理我们项目的依赖==。

```bash
go mod init 当前文件夹
```

<img src="go%E7%9A%84%E5%8C%85%E7%AE%A1%E7%90%86.assets/image-20220104233543610.png" alt="image-20220104233543610" style="zoom:50%;" />

### 2.2 引入其他项目的包

==注意==,我们在test下建了==一个`calc`目录==，并在里面建了一个文件`calc.go`

```go
package calc

// 自定义包，最好和文件夹统一起来

// 私有变量
var age = 10
// 公有变量
var Name = "张三"

// 首字母大写，表示共有方法
func Add(x, y int)int  {
	return x + y
}
func Sub(x, y int)int  {
	return x - y
}
```

在==当先项目test下==引用，只需要如下，==即路径要注意==

<font color=red>如何在非test项目下引用calc</font>

```go
package main
import (
	"fmt"
	"test/calc"
)
func main() {
	fmt.Printf("%v", calc.Add(2, 5))
}
```

### 2.3 go 自定义包

包（package）是多个Go源码的集合，==一个包可以简单理解为一个存放多个.go文件的文件夹==。该文件夹下面的所有go文件都要

在代码的第一行添加如下代码，声明该文件归属的包。

```go
package 包名
```

**注意事项**

- 一个文件夹下面直接包含的文件==只能归属一个package，同样一个package的文件不能在多个文件夹下==。
- 包名可以不和文件夹的名字一样，==包名不能包含-符号==。
- 包名为main的包为==应用程序的入口包==，这种包编译后会得到一个可执行文件，而==编译不包含main包的源代码则不会得到可执行文件==。

### 2.4 go中的init初始化函数

在Go 语言程序执行时==导入包语句会自动触发包内部init（）函数的调用==。需要注意的是：==`init()`函数没有参数也没有返回值。==

`init()`函数在程序运行时==自动被调用执行，不能在代码中主动调用它。==

==全局声明->init函数->main函数==

==包嵌套引用时的`init()`函数执行顺序与嵌套顺序相反==

## 3. go的第三方包

我们可以在 https://pkg.go.dev/ 查找看常见的golang第三方包

例如，前面找到前面我们需要下载的第三方包的地址

```
https://github.com/shopspring/decimal
```

然后安装这个包

- 方法1：`go get` 包全名 ==（全局）==

```bash
go get github.com/shopspring/decimal
```

- 方法2：`go mod download` ==（全局）==

```bash
go mod download
```

依赖包会自动下载到 `$GOPATH/pkg/mod`目录，并且==多个项目可以共享缓存的mod==，注意使用`go mod download`的时候，需要==首先在你的项目中引入第三方包==

- 方法3：`go mod vendor` 将依赖复制到当前项目的vendor（本项目）

```bash
go mod vendor
```

注意：使用go mod vendor的时候，==首先需要在你的项目里面引入第三方包==

## 4. go常见命令

- go download：下载依赖的module到本地cache
- go edit：编辑go.mod文件
- go graph：打印模块依赖图
- go init：在当前文件夹下初始化一个新的module，创建go.mod文件
- tidy：增加丢失的module，去掉未使用的module
- vendor：将依赖复制到vendor下
- verify：校验依赖，检查下载的第三方库有没有本地修改，如果有修改，则会返回非0，否则校验成功

## 5. 引入第三方包

同样在test目录下，修改`main.go`

```go
package main
import (
	"fmt"
	"github.com/shopspring/decimal" // 引入了第三方包
	"test/calc"
)
func main() {
	fmt.Printf("%v \n", calc.Add(2, 5))
	// 打印公有变量
	fmt.Println(calc.Name)

	_, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
}
```

引入后，我们==运行项目==，就会去下载了，下载完成后，则`go.mod`变为

```bash
module test

go 1.14

require github.com/shopspring/decimal v1.3.1 // indirect
```

同时我们生成了一个`go.sum`文件

```bash
github.com/shopspring/decimal v1.3.1 h1:2Usl1nmF/WZucqkFZhnfFYxxxu8LG21F6nPQBE5gKV8=
github.com/shopspring/decimal v1.3.1/go.mod h1:DKyhrW/HYNuLGql+MJL6WCR6knT2jwCFRcu2hWCYk4o=

```

<img src="go%E7%9A%84%E5%8C%85%E7%AE%A1%E7%90%86.assets/image-20220105001916630.png" alt="image-20220105001916630" style="zoom:67%;" />

