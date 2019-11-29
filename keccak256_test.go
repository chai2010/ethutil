// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"testing"
)

func TestKeccak256Hash(t *testing.T) {
	const s = "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
	tAssert(t, Keccak256Hash([]byte("")) == s)
}
