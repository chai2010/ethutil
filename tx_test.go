// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"strings"
	"testing"
)

func TestTxData_SigData(t *testing.T) {
	var tests = []struct {
		Name        string
		ChainId     int64
		BlockNumber int64
		HexSigData  string
		HexHash     string
		TxData
	}{
		// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md
		{
			"EIP155-Example",
			1, EIP155_FORK_BLKNUM,
			"0xec098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a764000080018080",
			"0xdaf5a779ae972f972197303d7b574746c7ef83eadac0f2791ad23db92e4c8e53",
			TxData{
				Nonce:    9,
				GasPrice: "20" + strings.Repeat("0", 9),
				GasLimit: "21000",
				To:       "0x3535353535353535353535353535353535353535",
				Value:    "1" + strings.Repeat("0", 18),
				Data:     nil,
			},
		},

		// 书中 RLP-Encoded 并不是用于签名的编码数据, 还包含来 V/R/S
		// https://github.com/ethereumbook/ethereumbook/blob/develop/06transactions.asciidoc
		// https://github.com/ethereumbook/ethereumbook/blob/develop/code/web3js/raw_tx/raw_tx_demo.js
		//
		// $ node raw_tx_demo.js
		// RLP-Encoded Tx: 0xe6808609184e72a0008303000094b0920c523d582040f2bcb1bd7fb1c7c1...
		// Tx Hash: 0xaa7f03f9f4e52fcf69f836a6d2bbc7706580adce0a068ff6525ba337218e6992
		// Signed Raw Transaction: 0xf866808609184e72a0008303000094b0920c523d582040f2bcb1...
		{
			"精通以太坊例子",

			// 只编码6个交易数据
			// 书中的RLP编码数据不是用于签名, 因此忽略该检查
			0, 0, "",

			"0xaa7f03f9f4e52fcf69f836a6d2bbc7706580adce0a068ff6525ba337218e6992",
			TxData{
				Nonce:    0,
				GasPrice: "0x09184e72a000",
				GasLimit: "0x30000",
				To:       "0xb0920c523d582040f2bcb1bd7fb1c7c1ecebdb34",
				Value:    "0x00",
				Data:     nil,
			},
		},
	}

	for _, tv := range tests {
		if tv.HexSigData != "" {
			var got = Hex("0x%x", tv.SigDataEx(tv.ChainId, tv.BlockNumber)).String()
			tAssert(t, got == tv.HexSigData, tv.Name, tv.HexSigData, got)
		}

		var h = Hex("0x%x", tv.HashEx(tv.ChainId, tv.BlockNumber)).String()
		tAssert(t, h == tv.HexHash, tv.Name, tv.HexHash, h)
	}
}
