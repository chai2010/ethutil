// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil_test

import (
	"fmt"
	"log"

	"github.com/chai2010/ethutil"
)

func ExampleKeyStoreDecryptKey() {
	const keyjson = `{
		"crypto" : {
			"cipher" : "aes-128-ctr",
			"cipherparams" : {
				"iv" : "83dbcc02d8ccb40e466191a123791e0e"
			},
			"ciphertext" : "d172bf743a674da9cdad04534d56926ef8358534d458fffccd4e6ad2fbde479c",
			"kdf" : "scrypt",
			"kdfparams" : {
				"dklen" : 32,
				"n" : 262144,
				"r" : 1,
				"p" : 8,
				"salt" : "ab0c7876052600dd703518d6fc3fe8984592145b591fc8fb5c6d43190334ba19"
			},
			"mac" : "2103ac29920d71da29f15d75b4a16dbe95cfd7ff8faea1056c33131d846e3097"
		},
		"id" : "3198bc9c-6672-5ab3-d995-4942343ae5b6",
		"version" : 3
	}`

	uuid, privateKey, err := ethutil.KeyStoreDecryptKey(
		[]byte(keyjson), "testpassword",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid)
	fmt.Println(privateKey)

	// Output:
	// 3198bc9c-6672-5ab3-d995-4942343ae5b6
	// 7a28b5ba57c53603b0b07b56bba752f7784bf506fa95edc395f5cf6c7514fe9d
}
