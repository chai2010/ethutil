// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"regexp"
	"strings"
	"unicode"
)

// 从私钥生成账户地址
func GenAddressFromPrivateKey(privateKey string) string {
	return GenAddressFromPublicKey(GenPublicKey(privateKey))
}

// 公钥生成账户地址
// 结尾的20个字节, 对应十六进制的40个字符
// 包含十六进制的 0x 开头
func GenAddressFromPublicKey(publicKey string) string {
	// 去掉公钥开头的 04 部分
	var xy = publicKey[len("04"):]

	// 转换为字节格式, 并计算 Keccak256 哈希
	var hash = Keccak256Hash(MustBigint(xy, 16).Bytes())

	// 取十六进制格式的最后40个字节作为地址
	return "0x" + hash[len(hash)-40:]
}

// 是否为有效地址格式, 开头的0x可选
func IsValidAddress(key string) bool {
	re := regexp.MustCompile("^(0[xX])?[0-9a-fA-F]{40}$")
	return re.MatchString(key)
}

// 生成EIP55校验的地址
//
// EIP55通过十六进制的大小写来保存校验码信息
//
// 参考 https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md
func GenEIP55Address(address string) string {
	// 去掉 0x 开头
	if len(address) > 2 {
		if address[0] == '0' && (address[1] == 'x' || address[1] == 'X') {
			address = address[len("0x"):]
		}
	}

	// 转换为小写十六进制
	address = strings.ToLower(address)

	// 计算 Keccak-256 哈希
	var h = Keccak256Hash([]byte(address))

	// EIP-55编码
	// 每个字符对应位置的字符小于8, 则保留小写
	s1 := []byte(address)
	s2 := []byte(h)
	for i := 0; i < len(s1); i++ {
		if s2[i] < '8' {
			s1[i] = s1[i]
		} else {
			s1[i] = byte(unicode.ToUpper(rune(s1[i])))
		}
	}

	// 得到新地址
	return "0x" + string(s1)
}

// 校验EIP55地址
//
// EIP55通过十六进制的大小写来保存校验码信息
//
// 参考 https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md
func IsValidEIP55Address(address string) bool {
	// 添加缺少的 0x 前缀
	if len(address) > 2 {
		if address[0] != '0' || (address[1] != 'x' && address[1] != 'X') {
			address = "0x" + address
		}
	}

	// 重新生成规范的EIP55格式地址
	var goden = GenEIP55Address(address)

	// 校验新生成的地址是否相等
	return goden == address
}
