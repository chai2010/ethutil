// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

package ethutil

import "testing"

func tAssert(tb testing.TB, ok bool, a ...interface{}) {
	if !ok {
		tb.Helper()
		tb.Fatal(a...)
	}
}

func tAssertf(tb testing.TB, ok bool, format string, a ...interface{}) {
	if !ok {
		tb.Helper()
		tb.Fatalf(format, a...)
	}
}
