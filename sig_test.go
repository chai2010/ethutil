// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import "testing"

func TestVerifySignature(t *testing.T) {
	sig := "0x638a54215d80a6713c8d523a6adc4e6e73652d859103a36b700851cb0e61b66b8ebfc1a610c57d732ec6e0a8f06a9a7a28df5051ece514702ff9cdff0b11f454"
	key := "0x03ca634cae0d49acb401d8a4c6b6fe8c55b70d115bf400769cc1400f3258cd3138"
	msg := "0xd301ce462d3e639518f482c7f03821fec1e602018630ce621e1e7851c12343a6"

	// func VerifySignature(pubkey, digestHash string, r, s *big.Int) bool {

	_ = sig
	_ = key
	_ = msg
}
