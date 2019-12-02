// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 生成以太坊私钥
// 十六进制格式, 不包含0x头
func GenPrivateKey() string {
	// 256bit 对应 32 字节
	var buf [32]byte

	// 生成256bit随机数
	// 必须是由加密库的强随机数函数生成!!!
	if _, err := rand.Read(buf[:]); err != nil {
		panic(err)
	}

	// 得到对应的256bit整数
	// 然后必须对 secp256k1 模素数取模(及小概率会超出, 那是无效的)
	var key = new(big.Int)
	key.SetBytes(buf[:]).Mod(key, _SECP256K1_P)

	// 最终以十六进制的格式输出
	// 256bit对应32字节, 对应64个十六进制字符
	return fmt.Sprintf("%064x", key)
}

// 生成公钥(04开头)
// 十六进制格式, 不包含0x头
func GenPublicKey(privateKey string) string {
	// 私钥展开为 big.Int
	var k = Hex(privateKey).MustBigint()

	// 生成公钥算法
	// secp256k1 椭圆曲线上定义的加法运算
	// 公钥 K = k*G, K 是k*G得到的椭圆上的点
	var Kx, Ky = S256().ScalarBaseMult(k.Bytes())

	// 格式化公钥
	// 以太坊公钥以04开头, 然后是x和y的十六进制格式字符串
	var publicKey = fmt.Sprintf("04%064x%64x", Kx, Ky)

	// OK
	return publicKey
}
