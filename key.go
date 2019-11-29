// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"

	"github.com/chai2010/ethutil/secp256k1"
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

// 是否为有效下私钥格式
// 私钥必须是十六进制格式, 开头的0x可选
// 没有检查超出素数P的情况
func IsValidPrivateKey(key string) bool {
	re := regexp.MustCompile("^(0[xX])?[0-9a-fA-F]{64}$")
	return re.MatchString(key)
}

// 生成公钥(04开头)
// 十六进制格式, 不包含0x头
func GenPublicKey(privateKey string) string {
	// 私钥展开为 big.Int
	var k = AsBigint(privateKey, 16)

	// 生成公钥算法
	// secp256k1 椭圆曲线上定义的加法运算
	// 公钥 K = k*G, K 是k*G得到的椭圆上的点
	var Kx, Ky = secp256k1.S256().ScalarBaseMult(k.Bytes())

	// 格式化公钥
	// 以太坊公钥以04开头, 然后是x和y的十六进制格式字符串
	var publicKey = fmt.Sprintf("04%064x%64x", Kx, Ky)

	// OK
	return publicKey
}

// 是否为有效下公钥格式
// 公钥必须是十六进制格式, 开头的0x可选
// 不计0x开头, 公钥的十六进制格式为130个字节
// 公钥开头的04表示未压缩点, 是以太坊唯一的格式
func IsValidPublicKey(publicKey string) bool {
	re := regexp.MustCompile("^(0[xX])?04[0-9a-fA-F]{128}$")
	return re.MatchString(publicKey)
}

// 公钥生成账户地址
// 结尾的20个字节, 对应十六进制的40个字符
// 包含十六进制的 0x 开头
func GenAddressFromPublicKey(publicKey string) string {
	// 去掉公钥开头的 04 部分
	var xy = publicKey[len("04"):]

	// 转换为字节格式, 并计算 Keccak256 哈希
	var hash = Keccak256Hash(AsBigint(xy, 16).Bytes())

	// 取十六进制格式的最后40个字节作为地址
	return "0x" + hash[len(hash)-40:]
}
