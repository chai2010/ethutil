// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

// https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses

// 比特币荣耀地址生成器
//
// 使用例子:
//
//	go run main.go
//	go run main.go -n=3
//	go run main.go -p=a
//	go run main.go -p=abc
//	go run main.go -p=a -s=bc
//	go run main.go -re="^a\d.*"
//
package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/chai2010/ethutil"
	"github.com/chai2010/ethutil/base58"
	"golang.org/x/crypto/ripemd160"
)

var (
	flagPrefix    = flag.String("p", "", "地址前缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagSuffix    = flag.String("s", "", "地址后缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagRegexp    = flag.String("re", "", "正则模式, 需要字节验证")
	flagPrefixPrv = flag.String("p-prv", "", "私钥前缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagNumKey    = flag.Int("n", 1, "生成几个地址")
	flagHelp      = flag.Bool("h", false, "显示帮助")
)

func init() {
	flag.Usage = func() {
		fmt.Println("比特币荣耀地址生成器")
		fmt.Println()

		fmt.Println("  ethutil-btc-vanity-gen            # 生成1个地址")
		fmt.Println("  ethutil-btc-vanity-gen -n=3       # 生成3个地址")
		fmt.Println("  ethutil-btc-vanity-gen -p=a -s=bc # ab开头, c结尾")
		fmt.Println()

		fmt.Println("参数说明:")
		flag.PrintDefaults()
		fmt.Println()

		fmt.Println("https://github.com/chai2010/ethutil")
		fmt.Println("版权 @2019 柴树杉")
	}
}

func main() {
	flag.Parse()

	if *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	if s := *flagPrefix; s != "" {
		if !regexp.MustCompile("^[0-9a-f]*$").MatchString(string(s)) {
			fmt.Printf("非法的前缀参数: %q\n", s)
			os.Exit(1)
		}
	}
	if s := *flagSuffix; s != "" {
		if !regexp.MustCompile("^[0-9a-f]*$").MatchString(string(s)) {
			fmt.Printf("非法的后缀参数: %q\n", s)
			os.Exit(1)
		}
	}
	if s := *flagRegexp; s != "" {
		if _, err := regexp.Compile(s); err != nil {
			fmt.Printf("非法的正则参数: %q\n", s)
			os.Exit(1)
		}
	}

	for i := 0; i < *flagNumKey; i++ {
		key, addr := genVanityBtc(*flagPrefixPrv, *flagPrefix, *flagSuffix, *flagRegexp)
		fmt.Println(i, addr, key)
	}
}

func genVanityBtc(prefixPrv, prefix, suffix, reExpr string) (key, addr string) {
	fmt.Printf("prefixPrv=%q, prefix=%q, suffix=%q, reExpr=%q\n", prefixPrv, prefix, suffix, reExpr)

	if reExpr == "" {
		reExpr = ".*"
	}
	var re = regexp.MustCompile(reExpr)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("key = %q, addr = %q\n", key, addr)
			panic(r)
		}
	}()

	for {
		k := NewBtcKey(true)
		key = k.KeyString()
		addr = k.AddressString()

		if strings.HasPrefix(key, prefixPrv) &&
			strings.HasPrefix(addr[2:], prefix) &&
			strings.HasSuffix(addr, suffix) &&
			re.MatchString(addr[2:]) {
			return key, addr
		}
	}
}

type BtcKey struct {
	Compressed bool
	D          *big.Int
}

func NewBtcKey(compressed bool) *BtcKey {
	key, err := ecdsa.GenerateKey(ethutil.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return &BtcKey{
		Compressed: compressed,
		D:          key.D,
	}
}

func (p *BtcKey) Key() *big.Int {
	return p.D
}

func (p *BtcKey) KeyBytes() []byte {
	if p.Compressed {
		return append(p.bigintPaddedBytes(p.D, 32), 0x01)
	} else {
		return p.bigintPaddedBytes(p.D, 32)
	}
}

func (p *BtcKey) KeyString() string {
	return base58.CheckEncode(p.KeyBytes(), 0x80)
}

func (p *BtcKey) PubKey() (x, y *big.Int) {
	var Kx, Ky = ethutil.S256().ScalarBaseMult(p.D.Bytes())
	return Kx, Ky
}
func (p *BtcKey) PubKeyBytes() []byte {
	if p.Compressed {
		if x, y := p.PubKey(); y.Bit(0) == 0 {
			return append([]byte{0x02}, p.bigintPaddedBytes(x, 32)...)

		} else {
			return append([]byte{0x03}, p.bigintPaddedBytes(x, 32)...)
		}
	} else {
		x, y := p.PubKey()
		pubkeyBytes := []byte{0x04}
		pubkeyBytes = append(pubkeyBytes, p.bigintPaddedBytes(x, 32)...)
		pubkeyBytes = append(pubkeyBytes, p.bigintPaddedBytes(y, 32)...)
		return pubkeyBytes
	}
}
func (p *BtcKey) PubKeyString() string {
	return fmt.Sprintf("%x", p.PubKeyBytes())
}

func (p *BtcKey) AddressString() string {
	var pubkeyBytes = p.PubKeyBytes()
	var addrBytes = p.ripemd160Hash(p.sha256Hash(pubkeyBytes))
	return base58.CheckEncode(addrBytes, 0x00)
}

func (p *BtcKey) String() string {
	return p.AddressString()
}

func (p *BtcKey) ripemd160Hash(data []byte) []byte {
	h := ripemd160.New()
	h.Reset()
	h.Write(data)
	return h.Sum(nil)
}

func (p *BtcKey) sha256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// PaddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
func (p *BtcKey) bigintPaddedBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	p.bigintReadBits(bigint, ret)
	return ret
}

// ReadBits encodes the absolute value of bigint as big-endian bytes. Callers must ensure
// that buf has enough space. If buf is too short the result will be incomplete.
func (p *BtcKey) bigintReadBits(bigint *big.Int, buf []byte) {
	const (
		// number of bits in a big.Word
		wordBits = 32 << (uint64(^big.Word(0)) >> 63)
		// number of bytes in a big.Word
		wordBytes = wordBits / 8
	)

	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			buf[i-1] = byte(d)
			d >>= 8
			i--
		}
	}
}
