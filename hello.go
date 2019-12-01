// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

// +build ignore

package main

import (
	"fmt"
	"log"
	"math/big"

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

	fmt.Println(ethutil.Hex("%x", genTxData1().Hash()))
}


// 精通以太坊例子
// Tx Hash: 0xaa7f03f9f4e52fcf69f836a6d2bbc7706580adce0a068ff6525ba337218e6992
func genTxData1() *ethutil.TxData {
	var tx = &ethutil.TxData{}
	{
		tx.Nonce = 0
		tx.GasPrice, _ = new(big.Int).SetString("0x09184e72a000"[2:], 16)
		tx.GasLimit = 0x30000

		tx.To = new([20]byte)
		to := ethutil.Hex("0xb0920c523d582040f2bcb1bd7fb1c7c1ecebdb34").MustBytes()
		copy(tx.To[:], to)

		tx.Value = big.NewInt(0)
	}
	return tx
}
