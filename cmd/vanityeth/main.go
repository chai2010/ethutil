// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

// 生成荣耀地址
//
// 使用例子:
//
//	go run main.go
//	go run main.go -n=3
//	go run main.go -p=a
//	go run main.go -p=abc
//	go run main.go -p=a -s=bc
//
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/chai2010/ethutil"
)

var (
	flagPrefix = flag.String("p", "", "地址前缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagSuffix = flag.String("s", "", "地址后缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagNumKey = flag.Int("n", 1, "生成几个地址")
)

func main() {
	flag.Parse()

	for i := 0; i < *flagNumKey; i++ {
		key, addr := genVanityEth(*flagPrefix, *flagSuffix)
		fmt.Println(i, addr, key)
	}
}

func genVanityEth(prefix, suffix string) (key, addr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("key = %q, addr = %q\n", key, addr)
			panic(r)
		}
	}()

	for {
		key = ethutil.GenPrivateKey()
		addr = ethutil.GenAddressFromPrivateKey(key)
		if strings.HasPrefix(addr[2:], prefix) && strings.HasSuffix(addr, suffix) {
			return
		}
	}
}
