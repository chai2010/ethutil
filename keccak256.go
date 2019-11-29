// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

// Keccak-256 哈希
// 十六进制格式, 不含0x前缀
func Keccak256Hash(data ...[]byte) string {
	keccak256 := sha3.NewLegacyKeccak256()
	for _, v := range data {
		keccak256.Write(v)
	}
	return fmt.Sprintf("%x", keccak256.Sum(nil))
}
