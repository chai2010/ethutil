// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"testing"
)

// 测试数据来自《精通以太坊》
// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc

func TestGenPublicKey(t *testing.T) {
	const privateKey = "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const publicKey = "046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0"

	tAssert(t, GenPublicKey(privateKey) == publicKey)
}
