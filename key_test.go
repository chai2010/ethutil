// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

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

	// 必须以0x开头, 否则解析数字时无法知道是否为十六进制格式
	const privateKey = "0x" + "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const publicKey = "0x" + "046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0"

	tAssert(t, GenPublicKey(privateKey) == publicKey)
}

func TestIsValidPublicKey(t *testing.T) {
	// 测试数据来自《精通以太坊》
	// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc

	// 必须以0x开头, 否则解析数字时无法知道是否为十六进制格式
	const privateKey = "0x" + "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const publicKey = "0x" + "046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0"

	tAssert(t, IsValidPublicKey(GenPublicKey(privateKey)))
	tAssert(t, IsValidPublicKey(publicKey))
}
