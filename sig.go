// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/chai2010/ethutil/btcec"
)

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

// 私钥签名
// 涉及的字符串都是十六进制编码格式
func Sign(privateKey, digestHash string) (hexSig string, err error) {
	var c elliptic.Curve = btcec.S256()
	var priv = new(btcec.PrivateKey)

	priv.PublicKey.Curve = c
	priv.D = Hex(privateKey).MustBigint()
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(Hex(privateKey).MustBytes())

	sig, err := btcec.SignCompact(btcec.S256(), priv, Hex(digestHash).MustBytes(), false)
	if err != nil {
		return "", err
	}

	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[64] = v

	return Hex("%x", sig).String(), nil
}

// 公钥验证签名
// 涉及的字符串都是十六进制编码格式
func VerifySignature(pubkey, hash, signature string) bool {
	if len(signature) != 128 {
		return false
	}
	sig := &btcec.Signature{
		R: Hex(signature[:64]).MustBigint(),
		S: Hex(signature[64:]).MustBigint(),
	}
	key, err := btcec.ParsePubKey(
		Hex(pubkey).MustBytes(), btcec.S256(),
	)
	if err != nil {
		return false
	}
	// Reject malleable signatures. libsecp256k1 does this check but btcec doesn't.
	if sig.S.Cmp(secp256k1halfN) > 0 {
		return false
	}
	return sig.Verify(Hex(hash).MustBytes(), key)
}

// 签名导出公钥
// 涉及的字符串都是十六进制编码格式
func SigToPub(msgHash, sig string) (publicKey string, err error) {
	pub, _, err := btcec.RecoverCompact(btcec.S256(),
		Hex(sig).MustBytes(),
		Hex(msgHash).MustBytes(),
	)

	publicKey = fmt.Sprintf("04%064x%64x", pub.X, pub.Y)
	return
}

// secp256k1 椭圆曲线
func S256() elliptic.Curve {
	return btcec.S256()
}
