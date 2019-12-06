// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"bytes"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/chai2010/ethutil/rlp"
)

// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md
// FORK_BLKNUM: 2,675,000
// CHAIN_ID: 1 (main net)

// EIP分叉时区块高度
// 之后交易的签名数据从6个变成9个
const EIP155_FORK_BLKNUM = 2675000

// 交易数据
type TxData struct {
	Nonce    uint64
	GasPrice string
	GasLimit string
	To       string // 0地址 表示合约
	Value    string
	Data     []byte

	V *big.Int
	R *big.Int
	S *big.Int
}

// RLP编码签名过后的交易
func (p *TxData) EncodeRLP() []byte {
	var buf bytes.Buffer
	rlp.Encode(&buf, p)
	return buf.Bytes()
}

// 生成最新规范的要签名编码数据
func (p *TxData) SigData() []byte {
	// 输出EIP55之后的格式, 不包含 chainId
	return p.sigData(0, EIP155_FORK_BLKNUM)
}

// 生成要签名编码数据(支持旧规范)
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md
func (p *TxData) SigDataEx(chainId, blockNumber int64) []byte {
	// 输出EIP55之后的格式, 包含 chainId
	return p.sigData(chainId, blockNumber)
}

// 计算交易的哈希(可用于后续的签名)
func (p *TxData) Hash() []byte {
	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write(p.SigData())
	return keccak256.Sum(nil)
}

// 计算交易的哈希
func (p *TxData) HashEx(chainId, blockNumber int64) []byte {
	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write(p.SigDataEx(chainId, blockNumber))
	return keccak256.Sum(nil)
}

// 输出用于签名hash的RLP编码数据
func (p *TxData) sigData(chainId, blockNumber int64) []byte {
	// 解码Hex格式地址
	// 如果是零地址, 用 nil 表示
	var p_To *[20]byte
	if !Hex(p.To).IsZero() {
		var address [20]byte
		copy(address[:], Hex(p.To).MustBytes())
		p_To = &address
	}

	// 初始规范只编码交易的6个信息
	if blockNumber < EIP155_FORK_BLKNUM {
		var txData = &struct {
			Nonce    uint64
			GasPrice *big.Int
			GasLimit uint64
			To       *[20]byte // nil 表示合约
			Value    *big.Int
			Data     []byte
		}{
			Nonce:    p.Nonce,
			GasPrice: MustBigint(p.GasPrice, 0),
			GasLimit: MustUint64(p.GasLimit, 0),
			To:       p_To,
			Value:    MustBigint(p.Value, 0),
			Data:     p.Data,
		}

		var buf bytes.Buffer
		rlp.Encode(&buf, txData)
		return buf.Bytes()
	}

	// EIP155之后增加 V/R/S 三个元素
	// V 可以包含 chainId, 但是依然兼容以前旧的风格
	// 因为增加来内容, 导致旧版本的客户端不再兼容
	var txData = &struct {
		Nonce    uint64
		GasPrice *big.Int
		GasLimit uint64
		To       *[20]byte // nil 表示合约
		Value    *big.Int
		Data     []byte

		V *big.Int
		R *big.Int
		S *big.Int
	}{
		Nonce:    p.Nonce,
		GasPrice: MustBigint(p.GasPrice, 16),
		GasLimit: MustUint64(p.GasLimit, 16),
		To:       p_To,
		Value:    MustBigint(p.Value, 16),
		Data:     p.Data,
	}

	// 包含 chainId 信息
	if chainId > 0 {
		// v = CHAIN_ID * 2 + 35 or v = CHAIN_ID * 2 + 36
		txData.V = big.NewInt(chainId)
	} else {
		txData.V = big.NewInt(27)
	}

	// RLP编码
	var buf bytes.Buffer
	rlp.Encode(&buf, txData)
	return buf.Bytes()
}
