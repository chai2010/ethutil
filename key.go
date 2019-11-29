// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 生成以太坊私钥
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
