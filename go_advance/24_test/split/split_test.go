package split

import (
	"reflect"
	"testing"
)

// 测试函数必须以Test开头，必须接收一个*testing.T类型的参数

// 1.基本测试
func TestSplit(t *testing.T) {

	got := Split("a:b:c", ":")

	want := []string{"a", "b", "c"}

	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}

}

// 2.测试用例2
// 可以在`go test`命令后添加`-run`参数，它对应一个正则表达式，
// 只有函数名匹配上的测试函数才会被`go test`命令执行。
// -v表示显示测试的详细命令
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

// 3.分组测试 多个测试用例
func TestGroupSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}

// 4. 使用t.Run()执行子测试
// go test -v -run=SplitRun/simple 命令可以执行具体TestSplitRun中的simple对应的子测试用例
func TestSplitRun(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		// 使用t.Run()执行子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}

	// go test -cover可以查看测试覆盖率
	// go test -cover -coverprofile=c.out

}

// 5. 基准测试
// 基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，
// 基准测试必须要执行b.N次，这样的测试才有对照性，
// b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。
// 基准测试并不会默认执行，需要增加-bench参数，
// 所以我们通过执行go test -bench=Split命令执行基准测试
// 其中`BenchmarkSplit-4`表示对Split函数进行基准测试，
// 数字`4`表示`GOMAXPROCS`的值，这个对于并发基准测试很重要。
// `10000000`和`203ns/op`表示每次调用`Split`函数耗时`203ns`，
// 这个结果是`10000000`次调用的平均值。
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

// 待续
