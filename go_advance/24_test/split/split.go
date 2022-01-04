package split

import "strings"

// 1. test相关说明
// 测试依赖go test命令
// 在包目录内，所有以_test.go为后缀名的源代码文件
// 都是go test测试的一部分， 不会被go build编译到最终的可执行文件中。

// 2. *_test.go文件中三种类型的函数
// |   类型   |         格式          |              作用              |
//| :------: | :-------------------: | :----------------------------: |
//| 测试函数 |   函数名前缀为Test    | 测试程序的一些逻辑行为是否正确 |
//| 基准函数 | 函数名前缀为Benchmark |         测试函数的性能         |
//| 示例函数 |  函数名前缀为Example  |       为文档提供示例文档       |

// `go test`命令会遍历所有的`*_test.go`文件中符合上述命名规则的函数，
// 然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，
// 最后清理测试中生成的临时文件。

// 用sep分割s

func Split(target, sep string) (result []string) {
	// 返回sep在target中第一次出现的位置
	i := strings.Index(target, sep)

	for i > -1 {
		// 给切片扩容操作 难道忘了？
		result = append(result, target[:i])
		// 这是go字符串的截取操作，但是截取的是从i+len(sep)后开始的
		target = target[i+len(sep):]
		// 然后求在下一段出现的位置  然后继续分割
		i = strings.Index(target, sep)

	}

	// 最后一段
	result = append(result, target)
	return

}
