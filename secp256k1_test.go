// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"math/big"
	"testing"
)

func TestSECP256K1_P(t *testing.T) {
	tAssert(t, SECP256K1_P0 == SECP256K1_P1)

	var p, _ = new(big.Int).SetString(SECP256K1_P2, 10)
	tAssert(t, p.String() == SECP256K1_P2)
}

func TestSecp256k1_Fx_Fy(t *testing.T) {
	// 测试数据来自《精通以太坊》
	// https://github.com/ethereumbook/ethereumbook/blob/develop/04keys-addresses.asciidoc
	x, _ := new(big.Int).SetString(
		"49790390825249384486033144355916864607616083520101638681403973749255924539515",
		10,
	)
	y, _ := new(big.Int).SetString(
		"59574132161899900045862086493921015780032175291755807399284007721050341297360",
		10,
	)

	z := Secp256k1_Fx_Fy(x, y)
	tAssert(t, z.BitLen() == 0)
}
