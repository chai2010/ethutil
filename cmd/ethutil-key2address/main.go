// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chai2010/ethutil"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Println("Usage: ethutil-key2address key.file")
		fmt.Println("       ethutil-key2address -h")
		fmt.Println()

		fmt.Println("https://github.com/chai2010/ethutil")
		fmt.Println("版权 @2019 柴树杉")

		os.Exit(0)
	}

	for i := 1; i < len(os.Args); i++ {
		// 读取key文件
		keydata, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "错误: 读文件 %q 失败: %v\n", os.Args[i], err)
			os.Exit(-1)
		}

		// 去掉开头和结尾的空白部分
		keydata = bytes.TrimSpace(keydata)

		// 输出对应的地址
		fmt.Println(ethutil.GenAddressFromPrivateKey(string(keydata)))
	}
}
