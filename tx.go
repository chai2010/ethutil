// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"bytes"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/chai2010/ethutil/rlp"
)

// 交易数据
type TxData struct {
	Nonce    uint64
	GasPrice *big.Int
	GasLimit uint64
	To       *[20]byte // nil 表示合约
	Value    *big.Int
	Data     []byte
}

// 计算交易的哈希(可用于后续的签名)
func (p *TxData) Hash() []byte {
	type txData struct {
		Nonce    uint64
		GasPrice *big.Int
		GasLimit uint64
		To       *[20]byte // nil 表示合约
		Value    *big.Int
		Data     []byte
	}
	var q = &txData{
		Nonce:    p.Nonce,
		GasPrice: p.GasPrice,
		GasLimit: p.GasLimit,
		To:       p.To,
		Value:    p.Value,
		Data:     p.Data,
	}

	var buf bytes.Buffer
	rlp.Encode(&buf, q)

	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write(buf.Bytes())
	return keccak256.Sum(nil)
}
