// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"testing"
)

func Test_privateKey(t *testing.T) {
	tAssert(t, IsValidPrivateKey(GenPrivateKey()))

	tAssert(t, !IsValidPrivateKey(""))
	tAssert(t, !IsValidPrivateKey("abc"))
	tAssert(t, !IsValidPrivateKey(SECP256K1_P2))
}
