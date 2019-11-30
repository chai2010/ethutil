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
func NewHexString(v interface{}) HexString {
	return HexString(fmt.Sprintf("%x", v))
}

// 是否为有效
func (s HexString) Valid() bool {
	re := regexp.MustCompile("^(0[xX])?[0-9a-fA-F]+$")
	return re.MatchString(string(s))
}

// 转换为字节类型
// 忽略 0x 开头部分
func (s HexString) Bytes() []byte {
	if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		v, _ := hex.DecodeString(string(s[2:]))
		return v
	} else {
		v, _ := hex.DecodeString(string(s))
		return v
	}
}

func (s HexString) String() string {
	return string(s)
}
