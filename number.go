// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
)

// 解析整数(失败时panic)
// 支持二进制/八进制/十进制/十六进制
func MustInt64(s string, base int) int64 {
	v, err := ParseInt64(s, base)
	if err != nil {
		log.Panic(err)
	}
	return v
}

// 解析整数
// 支持二进制/八进制/十进制/十六进制
func ParseInt64(s string, base int) (v int64, err error) {
	// 0b/0o/0x 开头
	if len(s) > 2 && s[0] == '0' {
		switch {
		case s[1] == 'b' || s[1] == 'B':
			if base != 0 && base != 2 {
				return 0, fmt.Errorf("ethutil.ParseInt64: invalid base %d", base)
			}
			return strconv.ParseInt(s[2:], 2, 64)
		case s[1] == 'o' || s[1] == 'O':
			if base != 0 && base != 8 {
				return 0, fmt.Errorf("ethutil.ParseInt64: invalid base %d", base)
			}
			return strconv.ParseInt(s[2:], 8, 64)
		case s[1] == 'x' || s[1] == 'X':
			if base != 0 && base != 16 {
				return 0, fmt.Errorf("ethutil.ParseInt64: invalid base %d", base)
			}
			return strconv.ParseInt(s[2:], 16, 64)
		}
	}

	// 其它根据base解析(默认为十进制)
	if base > 0 {
		return strconv.ParseInt(s, base, 64)
	} else {
		return strconv.ParseInt(s, 10, 64)
	}
}

// 解析无符号整数(失败时panic)
// 支持二进制/八进制/十进制/十六进制
func MustUint64(s string, base int) uint64 {
	v, err := ParseUint64(s, base)
	if err != nil {
		log.Panic(err)
	}
	return v
}

// 解析无符号整数
// 支持二进制/八进制/十进制/十六进制
func ParseUint64(s string, base int) (v uint64, err error) {
	// 0b/0o/0x 开头
	if len(s) > 2 && s[0] == '0' {
		switch {
		case s[1] == 'b' || s[1] == 'B':
			if base != 0 && base != 2 {
				return 0, fmt.Errorf("ethutil.ParseUint64: invalid base %d", base)
			}
			return strconv.ParseUint(s[2:], 2, 64)
		case s[1] == 'o' || s[1] == 'O':
			if base != 0 && base != 8 {
				return 0, fmt.Errorf("ethutil.ParseUint64: invalid base %d", base)
			}
			return strconv.ParseUint(s[2:], 8, 64)
		case s[1] == 'x' || s[1] == 'X':
			if base != 0 && base != 16 {
				return 0, fmt.Errorf("ethutil.ParseUint64: invalid base %d", base)
			}
			return strconv.ParseUint(s[2:], 16, 64)
		}
	}

	// 其它根据base解析(默认为十进制)
	if base > 0 {
		return strconv.ParseUint(s, base, 64)
	} else {
		return strconv.ParseUint(s, 10, 64)
	}
}

// 解析大整数(失败时panic)
// 支持二进制/八进制/十进制/十六进制
func MustBigint(s string, base int) *big.Int {
	v, err := ParseBigint(s, base)
	if err != nil {
		log.Panic(err)
	}
	return v
}

// 解析大整数
// 支持二进制/八进制/十进制/十六进制
func ParseBigint(s string, base int) (v *big.Int, err error) {
	// 0b/0o/0x 开头
	if len(s) > 2 && s[0] == '0' {
		switch {
		case s[1] == 'b' || s[1] == 'B':
			if base != 0 && base != 2 {
				println("ParseBigint 11")
				return nil, fmt.Errorf("ethutil.ParseBigint: invalid base %d", base)
			}
			v, ok := new(big.Int).SetString(s[2:], 2)
			if !ok {
				println("ParseBigint 22")
				return nil, fmt.Errorf("ethutil.ParseBigint: invalid bigint %q", s)
			}
			return v, nil
		case s[1] == 'o' || s[1] == 'O':
			if base != 0 && base != 8 {
				println("ParseBigint 33")
				return nil, fmt.Errorf("ethutil.ParseBigint: invalid base %d", base)
			}
			v, ok := new(big.Int).SetString(s[2:], 8)
			if !ok {
				println("ParseBigint 44")
				return nil, fmt.Errorf("ethutil.ParseBigint: invalid bigint %q", s)
			}
			return v, nil
		case s[1] == 'x' || s[1] == 'X':
			if base != 0 && base != 16 {
				println("ParseBigint 55")
				return nil, fmt.Errorf("ethutil.ParseBigint: invalid base %d", base)
			}
			v, ok := new(big.Int).SetString(s[2:], 16)
			if !ok {
				println("ParseBigint 66")
				return nil, fmt.Errorf("ethutil.ParseBigint: invalid bigint %q", s)
			}
			return v, nil
		}
	}

	// 其它根据base解析(默认为十进制)
	if base > 0 {
		v, ok := new(big.Int).SetString(s, base)
		if !ok {
			println("ParseBigint 77")
			return nil, fmt.Errorf("ethutil.ParseBigint: invalid bigint %q", s)
		}
		return v, nil
	} else {
		v, ok := new(big.Int).SetString(s, 10)
		if !ok {
			println("ParseBigint 88")
			return nil, fmt.Errorf("ethutil.ParseBigint: invalid bigint %q", s)
		}
		return v, nil
	}
}
