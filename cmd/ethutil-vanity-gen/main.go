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
	"os"
	"regexp"
	"strings"

	"github.com/chai2010/ethutil"
)

var (
	flagPrefix = flag.String("p", "", "地址前缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagSuffix = flag.String("s", "", "地址后缀, 必须是十六进制格式 ([0-9a-f]*)")
	flagNumKey = flag.Int("n", 1, "生成几个地址")
	flagHelp   = flag.Bool("h", false, "显示帮助")
)

func init() {
	flag.Usage = func() {
		fmt.Println("以太坊荣耀地址生成器")
		fmt.Println()

		fmt.Println("  ethutil-vanity-gen            # 生成1个地址")
		fmt.Println("  ethutil-vanity-gen -n=3       # 生成3个地址")
		fmt.Println("  ethutil-vanity-gen -p=a -s=bc # ab开头, c结尾")
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
