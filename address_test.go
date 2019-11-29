// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import "testing"

import "strings"

// 测试数据来自《精通以太坊》
// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc

func TestGenAddressFromPrivateKey(t *testing.T) {
	const privateKey = "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const addresses = "0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9"

	var s = GenAddressFromPrivateKey(privateKey)
	tAssert(t, s == addresses, s)
}

func TestGenAddressFromPublicKey(t *testing.T) {
	const privateKey = "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	const addresses = "0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9"

	var s = GenAddressFromPublicKey(GenPublicKey(privateKey))
	tAssert(t, s == addresses, s)
}

func TestGenEIP55Address(t *testing.T) {
	for i, s := range []string{
		// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc
		"0x001d3F1ef827552Ae1114027BD3ECF1f086bA0F9",

		// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md
		"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
		"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
	} {
		var got = GenEIP55Address(strings.ToLower(s))
		tAssert(t, got == s, i, got, s)
	}
}

func TestCheckEIP55Address(t *testing.T) {
	tAssert(t, CheckEIP55Address("001d3F1ef827552Ae1114027BD3ECF1f086bA0F9"))
	tAssert(t, CheckEIP55Address("0x001d3F1ef827552Ae1114027BD3ECF1f086bA0F9"))

	tAssert(t, !CheckEIP55Address("001d3f1ef827552ae1114027bd3ecf1f086ba0f9"))
	tAssert(t, !CheckEIP55Address("0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9"))
}
