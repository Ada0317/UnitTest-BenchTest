package UnitTest_BenchTest

import (
	"strings"
	"testing"
	"time"
)

/*
基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：
func BenchmarkName(b *testing.B){
    // ...
}

基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，
基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。

*/
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("古藤老树昏鸦", "老")
	}
}

/*基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试
输出结果如下：
UnitTest-BenchTest $ go test -bench=Split
    goos: darwin
    goarch: amd64
    pkg: studytest
    BenchmarkSplit-8        15680685                74.62 ns/op
    PASS
    ok      studytest       2.255s


* 其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
* 10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。
*/

func BenchmarkTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer := time.NewTimer(1)
		<-timer.C
		timer.Reset(1)
	}
}

/*
我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。
其中，200 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配。
*/

/*我们将我们的Split函数优化如下: */
func SplitUp(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func BenchmarkSplitUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitUp("古藤老树昏鸦", "老")
	}
}

/*
这一次我们提前使用make函数将result初始化为一个容量足够大的切片，
而不再像之前一样通过调用append函数来追加。我们来看一下这个改进会带来多大的性能提升:
*/
