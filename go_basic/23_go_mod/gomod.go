package main

// 从Go 1.11起，如果目录在$GOPATH/src之外，而且当前的路径或是任何父目录包含go.mod文件时，都可以使用Go Module的命令。
// 相反如果是在$GOPATH/src之中，为了兼容，即使存在go.mod文件，也只能使用原有的GOPATH下的命令。
// 但是从Go 1.13开始，Go Module模式将成为开发的默认模式。

//go mod init 创建一个新Modules，初始化go.mod文件。
//go build, go test 构建命令，添加所需要的依赖包，同时写入go.mod文件。
//go list -m all 打印当前Modules的依赖包。
//go get 更改所需依赖包的版本(或添加新的依赖包)。
//go mod tidy 移除不需要的依赖包。

func main() {

}
