// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/chai2010/ethutil"
)

func main() {
	var password = "testpassword"
	var uuid = "3198bc9c-6672-5ab3-d995-4942343ae5b6"
	var privateKey = "7a28b5ba57c53603b0b07b56bba752f7784bf506fa95edc395f5cf6c7514fe9d"
	keyjson, err := ethutil.KeyStoreEncryptKey(uuid, privateKey, password)
	if err != nil {
		log.Fatal(err)
	}

	uuid, privateKey, err = ethutil.KeyStoreDecryptKey(
		keyjson, password,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid)
	fmt.Println(privateKey)
	fmt.Println(string(keyjson))
}
