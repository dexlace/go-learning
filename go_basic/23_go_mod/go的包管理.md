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

## 6.  godep

Go语言从v1.5开始开始引入`vendor`模式，如果项目目录下有vendor目录，那么go工具链会优先使用`vendor`内的包进行编译、测试等。

`godep`是一个通过vender模式实现的Go语言的第三方依赖管理工具，类似的还有由社区维护准官方包管理工具`dep`。

### 安装

执行以下命令安装`godep`工具。

```go
go get github.com/tools/godep
```

### 基本命令

安装好godep之后，在终端输入`godep`查看支持的所有命令。

```bash
godep save     将依赖项输出并复制到Godeps.json文件中
godep go       使用保存的依赖项运行go工具
godep get      下载并安装具有指定依赖项的包
godep path     打印依赖的GOPATH路径
godep restore  在GOPATH中拉取依赖的版本
godep update   更新选定的包或go版本
godep diff     显示当前和以前保存的依赖项集之间的差异
godep version  查看版本信息
```

使用`godep help [command]`可以看看具体命令的帮助信息。

### 使用godep

==在项目目录下执行`godep save`命令，会在当前项目中创建`Godeps`和`vender`两个文件夹。==

其中`Godeps`文件夹下有一个==`Godeps.json`的文件，==里面记录了项目所依赖的包信息。==`vender`文件夹下是项目依赖的包的源代码文件==

### vender机制

Go1.5版本之后开始支持，能够控制Go语言程序编译时依赖包==搜索路径的优先级==。

例如查找项目的某个依赖包，首先会在==项目根目录下的`vender`文件夹中查找，如果没有找到就会去`$GOAPTH/src`目录下查找==。

### godep开发流程

1. 保证程序能够正常编译
2. 执行`godep save`保存当前项目的所有第三方依赖的版本信息和代码
3. 提交Godeps目录和vender目录到代码库。
4. 如果要更新依赖的版本，==可以直接修改`Godeps.json`文件中的对应项==

## 7. go module

`go module`是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，`go module`将是Go语言默认的依赖管理工具。

### GO111MODULE

要启用`go module`支持首先要设置环境变量`GO111MODULE`，通过它可以开启或关闭模块支持，它有三个可选值：`off`、`on`、`auto`，默认值是`auto`。

1. `GO111MODULE=off`禁用模块支持，编译时会从`GOPATH`和`vendor`文件夹中查找包。
2. `GO111MODULE=on`启用模块支持，编译时会忽略`GOPATH`和`vendor`文件夹，只根据 `go.mod`下载依赖。
3. `GO111MODULE=auto`，==当项目在`$GOPATH/src`外且项目根目录有`go.mod`文件时，开启模块支持。==

==简单来说，设置`GO111MODULE=on`之后就可以使用`go module`了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。==

使用 go module 管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`。

### GOPROXY

Go1.11之后设置GOPROXY命令为：

```bash
export GOPROXY=https://goproxy.cn
```

Go1.13之后`GOPROXY`默认值为`https://proxy.golang.org`，在国内是无法访问的，所以十分建议大家设置GOPROXY，==这里我推荐使用[goproxy.cn](https://studygolang.com/topics/10014)。==

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

### go mod命令

常用的`go mod`命令如下：

```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

### go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```sh
module github.com/Q1mi/studygo/blogger

go 1.12

require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect
)
```

其中，

- `module`用来定义包名
- `require`用来定义依赖包及版本
- `indirect`表示间接引用

#### 依赖的版本

go mod支持语义化版本号，比如`go get foo@v1.2.3`，也可以跟git的分支或tag，比如`go get foo@master`，当然也可以跟git提交哈希，比如`go get foo@e3702bed2`。关于依赖的版本支持以下几种格式：

```go
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/vmihailenco/msgpack.v2 v2.9.1
gopkg.in/yaml.v2 <=v2.2.1
github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
latest
```

#### replace

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

```go
replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

### go get

在项目中执行`go get`命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行`go get -u=patch`将会升级到最新的修订版本
3. 运行`go get package@version`将会升级到指定的版本号version

如果下载所有依赖可以使用`go mod download`命令。

`go get` 以后我发现下载的包不在`src`目录下生成,而==全部到了`$GOPATH$/pkg`目录下==

### 整理依赖

我们在代码中删除依赖代码后，相关的依赖库并不会在`go.mod`文件中自动移除。这种情况下我们可以使用==`go mod tidy`命令更新`go.mod`中的依赖关系。==

### go mod edit

#### 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

```bash
go mod edit -fmt
```

#### 添加依赖项

```bash
go mod edit -require=golang.org/x/text
```

#### 移除依赖项

如果只是想修改`go.mod`文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在`go.mod`中移除`golang.org/x/text`包，可以使用如下命令：

```bash
go mod edit -droprequire=golang.org/x/text
```

关于`go mod edit`的更多用法可以通过`go help mod edit`查看。

## 在项目中使用go module

### 既有项目

如果需要对一个已经存在的项目启用`go module`，可以按照以下步骤操作：

1. 在项目目录下执行`go mod init`，生成一个`go.mod`文件。
2. 执行`go get`，查找并记录当前项目的依赖，同时生成一个`go.sum`记录每个依赖库的版本和哈希值。

### 新项目

对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

1. 执行`go mod init 项目名`命令，在当前项目文件夹下创建一个`go.mod`文件。
2. 手动编辑`go.mod`中的require依赖项或执行`go get`自动发现、维护依赖。

