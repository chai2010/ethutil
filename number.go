// 以太坊工具箱(零依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"fmt"
	"math/big"
	"strconv"
)

// 解析整数
// 支持二进制/八进制/十进制/十六进制
// 失败时可以同提供的默认值代替, 否则panic
func AsInt64(s string, defaultValue ...int64) int64 {
	v, err := ParseInt64(s)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(err)
	}
	return v
}

// 解析整数
// 支持二进制/八进制/十进制/十六进制
func ParseInt64(s string) (v int64, err error) {
	// 0b/0o/0x 开头
	if len(s) > 2 && s[0] == '0' {
		switch {
		case s[1] == 'b' || s[1] == 'B':
			return strconv.ParseInt(s[2:], 2, 64)
		case s[1] == 'o' || s[1] == 'O':
			return strconv.ParseInt(s[2:], 8, 64)
		case s[1] == 'x' || s[1] == 'X':
			return strconv.ParseInt(s[2:], 16, 64)
		}
	}

	// 其它解析为普通的十进制
	return strconv.ParseInt(s, 10, 64)
}

// 解析无符号整数
// 支持二进制/八进制/十进制/十六进制
// 失败时可以同提供的默认值代替, 否则panic
func AsUint64(s string, defaultValue ...uint64) uint64 {
	v, err := ParseUint64(s)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(err)
	}
	return v
}

// 解析无符号整数
// 支持二进制/八进制/十进制/十六进制
func ParseUint64(s string) (v uint64, err error) {
	// 0b/0o/0x 开头
	if len(s) > 2 && s[0] == '0' {
		switch {
		case s[1] == 'b' || s[1] == 'B':
			return strconv.ParseUint(s[2:], 2, 64)
		case s[1] == 'o' || s[1] == 'O':
			return strconv.ParseUint(s[2:], 8, 64)
		case s[1] == 'x' || s[1] == 'X':
			return strconv.ParseUint(s[2:], 16, 64)
		}
	}

	// 其它解析为普通的十进制
	return strconv.ParseUint(s, 10, 64)
}

// 解析大整数
// 支持二进制/八进制/十进制/十六进制
// 失败时可以同提供的默认值代替, 否则panic
func AsBigint(s string, defaultValue ...*big.Int) *big.Int {
	v, err := ParseBigint(s)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(err)
	}
	return v
}

// 解析大整数
// 支持二进制/八进制/十进制/十六进制
func ParseBigint(s string) (v *big.Int, err error) {
	// 0b/0o/0x 开头
	if len(s) > 2 && s[0] == '0' {
		switch {
		case s[1] == 'b' || s[1] == 'B':
			v, ok := new(big.Int).SetString(s[2:], 2)
			if !ok {
				return nil, fmt.Errorf("ethrpc.parseBigint: invalid bigint %q", s)
			}
			return v, nil
		case s[1] == 'o' || s[1] == 'O':
			v, ok := new(big.Int).SetString(s[2:], 8)
			if !ok {
				return nil, fmt.Errorf("ethrpc.parseBigint: invalid bigint %q", s)
			}
			return v, nil
		case s[1] == 'x' || s[1] == 'X':
			v, ok := new(big.Int).SetString(s[2:], 16)
			if !ok {
				return nil, fmt.Errorf("ethrpc.parseBigint: invalid bigint %q", s)
			}
			return v, nil
		}
	}

	// 其它解析为普通的十进制
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("ethrpc.parseBigint: invalid bigint %q", s)
	}
	return v, nil
}
