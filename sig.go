// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md

// 私钥签名
// 涉及的字符串都是十六进制编码格式
func Sign(privateKey, digestHash string) (sig string, err error) {
	panic("todo")
}

// 公钥验证签名
// 涉及的字符串都是十六进制编码格式
func VerifySignature(pubkey, digestHash, signature string) bool {
	panic("todo")
}

// 签名导出公钥
// 涉及的字符串都是十六进制编码格式
func GetPublicFromSign(msgHash, v, r, s string) (publicKey string, err error) {
	panic("todo")
}
