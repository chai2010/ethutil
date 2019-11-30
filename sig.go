// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

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
func GetPublicFromSign(signature string) (publicKey string, err error) {
	// 首先解码出 v/r/s 参数
	panic("todo")
}
