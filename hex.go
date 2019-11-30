// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

// 十六进制字符串
type HexString string

// 生成十六进制字符串
// 类型 fmt.Sprintf 函数, 但是只有一个参数时只做转型操作
func Hex(format string, a ...interface{}) HexString {
	if len(a) > 0 {
		return HexString(fmt.Sprintf(format, a...))
	}
	return HexString(format)
}

// 是否为有效
func (s HexString) IsValid() bool {
	re := regexp.MustCompile("^(0[xX])?[0-9a-fA-F]*$")
	return re.MatchString(string(s))
}

// 是否为以太坊地址
func (s HexString) IsValidAddress() bool {
	re := regexp.MustCompile("^(0[xX])?[0-9a-fA-F]{40}$")
	return re.MatchString(string(s))
}

// 是否为有效下私钥格式
// 私钥必须是十六进制格式, 开头的0x可选
// 没有检查超出素数P的情况
func (s HexString) IsValidPrivateKey() bool {
	re := regexp.MustCompile("^(0[xX])?[0-9a-fA-F]{64}$")
	return re.MatchString(string(s))
}

// 是否为有效下公钥格式
// 公钥必须是十六进制格式, 开头的0x可选
// 不计0x开头, 公钥的十六进制格式为130个字节
// 公钥开头的04表示未压缩点, 是以太坊唯一的格式
func (s HexString) IsValidPublicKey() bool {
	re := regexp.MustCompile("^(0[xX])?04[0-9a-fA-F]{128}$")
	return re.MatchString(string(s))
}

// 转换为字节类型
// 忽略 0x 开头部分, 返回错误
func (s HexString) ToBytes() ([]byte, error) {
	if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		return hex.DecodeString(string(s[2:]))
	} else {
		return hex.DecodeString(string(s))
	}
}

// 转换为字节类型
// 忽略 0x 开头部分
func (s HexString) Bytes() []byte {
	v, _ := s.ToBytes()
	return v
}

// 转型为字符串
func (s HexString) String() string {
	return string(s)
}
