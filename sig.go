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
func Sign(privateKey string, hash []byte) (sig []byte, err error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}

	var c elliptic.Curve = btcec.S256()
	var priv = new(btcec.PrivateKey)

	priv.PublicKey.Curve = c
	priv.D = Hex(privateKey).MustBigint()
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(Hex(privateKey).MustBytes())

	sig, err = btcec.SignCompact(btcec.S256(), priv, hash, false)
	if err != nil {
		return nil, err
	}

	// r: sig[0:32]
	// s: sig[32:64]
	// v: sig[64]
	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[64] = v

	// ok
	return sig, nil
}

// 公钥验证签名
// 涉及的字符串都是十六进制编码格式
func VerifySignature(pubkey string, hash, signature []byte) bool {
	if len(signature) < 64 {
		return false
	}
	sig := &btcec.Signature{
		R: new(big.Int).SetBytes(signature[:32]),
		S: new(big.Int).SetBytes(signature[32:64]),
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
	return sig.Verify(hash, key)
}

// 签名导出公钥
func SigToPub(hash, sig []byte) (publicKey string, err error) {
	// 将v移到开头, 并加上27, 对应比特币的签名格式
	btcsig := make([]byte, 64+1)
	btcsig[0] = sig[64] + 27
	copy(btcsig[1:], sig)

	pub, _, err := btcec.RecoverCompact(btcec.S256(), btcsig, hash)
	if err != nil {
		return "", err
	}

	publicKey = fmt.Sprintf("04%064x%064x", pub.X, pub.Y)
	return
}

// secp256k1 椭圆曲线
func S256() elliptic.Curve {
	return btcec.S256()
}
