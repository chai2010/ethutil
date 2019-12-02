// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

// +build ignore

package main

import (
	"fmt"
	"log"
	"strings"

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
	fmt.Println(ethutil.Hex("%x", genTxData2().Hash()))

	var sgiData = "0xec098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a764000080018080"
	fmt.Println(ethutil.Keccak256Hash(ethutil.Hex(sgiData).MustBytes()))

	// 0xec098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a764000080018080
	fmt.Println(ethutil.Hex("%x", genTxData2().SigData()))
	//fmt.Println(ethutil.Hex("%x", genTxData2().SigDataEIP155()))

}


// 精通以太坊例子
// Tx Hash: 0xaa7f03f9f4e52fcf69f836a6d2bbc7706580adce0a068ff6525ba337218e6992
func genTxData1() *ethutil.TxData {
	var tx = &ethutil.TxData{}
	{
		tx.Nonce = 0
		tx.GasPrice = "0x09184e72a000"
		tx.GasLimit = "0x30000"
		tx.To = "0xb0920c523d582040f2bcb1bd7fb1c7c1ecebdb34"
		tx.Value = "0"
	}
	return tx
}

// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md
func genTxData2() *ethutil.TxData {
	var tx = &ethutil.TxData{}
	{
		tx.Nonce = 9
		tx.GasPrice = "20"+strings.Repeat("0", 9)
		tx.GasLimit = "21000"
		tx.To = "0x3535353535353535353535353535353535353535"
		tx.Value = "1"+strings.Repeat("0", 18)
	}
	return tx
}

/*
Consider a transaction with nonce = 9, gasprice = 20 * 10**9, startgas = 21000, to = 0x3535353535353535353535353535353535353535, value = 10**18, data=''
*/