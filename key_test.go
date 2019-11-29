// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"testing"
)

func TestGenPrivateKey(t *testing.T) {
	tAssert(t, IsValidPrivateKey(GenPrivateKey()))
}

func TestIsValidPrivateKey(t *testing.T) {
	tAssert(t, IsValidPrivateKey(GenPrivateKey()))

	tAssert(t, !IsValidPrivateKey(""))
	tAssert(t, !IsValidPrivateKey("abc"))
	tAssert(t, !IsValidPrivateKey(SECP256K1_P2))
}

func TestGenPublicKey(t *testing.T) {
	// 测试数据来自《精通以太坊》
	// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc

	const privateKey = "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const publicKey = "046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0"

	tAssert(t, GenPublicKey(privateKey) == publicKey)
}

func TestIsValidPublicKey(t *testing.T) {
	// 测试数据来自《精通以太坊》
	// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc

	const privateKey = "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const publicKey = "046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0"

	tAssert(t, IsValidPublicKey(GenPublicKey(privateKey)))
	tAssert(t, IsValidPublicKey(publicKey))
}

func TestGenAddressFromPublicKey(t *testing.T) {
	// 测试数据来自《精通以太坊》
	// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc

	const privateKey = "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const addresses = "0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9"

	var s = GenAddressFromPublicKey(GenPublicKey(privateKey))
	tAssert(t, s == addresses, s)
}
