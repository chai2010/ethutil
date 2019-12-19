// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

// 以太坊KeyStore文件查看器
//
// 创建KeyStore文件步骤:
// 1. 准备好私钥, 保存到key.txt文件
// 2. 执行 geth account import --keystore `pwd` key.txt
// 3. 通过步骤2中输入的密码可以查看生成的文件
// 4. 私钥可以使用 <ethutil-vanity-gen> 工具生成
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chai2010/ethutil"
)

var (
	flagPassword  = flag.String("password", "", "设置KeyStore文件的密码(默认手工输入)")
	flagGlobMatch = flag.Bool("glob-match", false, "采用Glob模式匹配文件名")
	flagShowKey   = flag.Bool("show-key", false, "显示私钥")
	flagHelp      = flag.Bool("h", false, "显示帮助")
)

func init() {
	flag.Usage = func() {
		fmt.Println("以太坊KeyStore文件查看器")
		fmt.Println()

		fmt.Println("  # 指定文件名")
		fmt.Println("  ethutil-cat-keystore file.json")
		fmt.Println("  ethutil-cat-keystore file.json file2.json")
		fmt.Println()

		fmt.Println("  # 显示私钥")
		fmt.Println("  ethutil-cat-keystore -show-key file.json")
		fmt.Println("  ethutil-cat-keystore -show-key file.json file2.json")
		fmt.Println()

		fmt.Println("  # Glob模式匹配")
		fmt.Println("  ethutil-cat-keystore -glob-match file?.json")
		fmt.Println("  ethutil-cat-keystore -glob-match file*.json")
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

	// 显示帮助信息
	if *flagHelp || len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	// 指定输入文件名
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "错误: 缺少文件名!\n")
		os.Exit(-1)
	}

	// 读取文件名列表
	var files = flag.Args()

	// Glob模式展开文件列表
	if *flagGlobMatch {
		files = []string{}
		for _, pattern := range flag.Args() {
			matches, err := filepath.Glob(pattern)
			if err != nil {
				fmt.Fprintf(os.Stderr, "错误: %v\n", err)
				os.Exit(-1)
			}
			files = append(files, matches...)
		}
	}

	// 手工输入密码
	if *flagPassword == "" {
		fmt.Printf("输入密码: ")
		*flagPassword = string(GetPasswdMasked())
		if len(*flagPassword) <= 0 {
			fmt.Fprintf(os.Stderr, "错误: 密码不能为空!\n")
			os.Exit(-1)
		}
	}

	// 处理全部的ks文件
	for _, s := range files {
		// 读取ks文件
		keyjson, err := ioutil.ReadFile(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "错误: 读文件 %q 失败: %v\n", s, err)
			os.Exit(-1)
		}

		// 解码ks文件
		uuid, privateKey, err := ethutil.KeyStoreDecryptKey(
			[]byte(keyjson), *flagPassword,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "错误: 解密 %q 失败: %v\n", s, err)
			os.Exit(-1)
		}

		if *flagShowKey {
			fmt.Printf("<uuid:%s> <address:%s> <key:%s> <file:%s>\n", uuid,
				ethutil.GenAddressFromPrivateKey(privateKey), privateKey,
				filepath.Clean(s),
			)
		} else {
			fmt.Printf("<uuid:%s> <address:%s> <file:%s>\n", uuid,
				ethutil.GenAddressFromPrivateKey(privateKey),
				filepath.Clean(s),
			)
		}
	}
}
