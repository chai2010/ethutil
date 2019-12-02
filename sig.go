// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/chai2010/ethutil/secp256k1"
)

// 私钥签名
// 涉及的字符串都是十六进制编码格式
func Sign(privateKey, digestHash string) (r, s *big.Int, err error) {
	var key = new(ecdsa.PrivateKey)

	key.PublicKey.Curve = secp256k1.S256()
	key.D = Hex(privateKey).MustBigint()
	key.PublicKey.X, key.PublicKey.Y = key.PublicKey.Curve.ScalarBaseMult(Hex(privateKey).MustBytes())

	r, s, err = ecdsa.Sign(rand.Reader, key, Hex(digestHash).MustBytes())
	return
}

// 公钥验证签名
// 涉及的字符串都是十六进制编码格式
func VerifySignature(pubkey, digestHash string, r, s *big.Int) bool {
	if len(pubkey) != 2+64+64 {
		log.Panicf("invalid pubkey: %q", pubkey)
	}

	var key = new(ecdsa.PublicKey)

	key.Curve = secp256k1.S256()
	key.X = MustBigint(pubkey[2:][:64], 16)
	key.Y = MustBigint(pubkey[2+64:], 16)

	return ecdsa.Verify(key, Hex(digestHash).MustBytes(), r, s)
}

// 签名导出公钥
// 涉及的字符串都是十六进制编码格式
func GetPublicFromSign(msgHash string, v int, r, s *big.Int) (publicKey string, err error) {
	key, err := recoverKeyFromSignature(
		r, s, Hex(msgHash).MustBytes(), v,
		false,
	)
	if err != nil {
		return "", err
	}

	publicKey = fmt.Sprintf("04%064x%64x", key.X, key.Y)
	return
}

// 恢复公钥
func recoverKeyFromSignature(
	sig_R, sig_S *big.Int,
	msg []byte,
	iter int, doChecks bool,
) (*ecdsa.PublicKey, error) {
	// 1.1 x = (n * i) + r
	curve := secp256k1.S256()

	// 后面的乘法好像没有意义
	Rx := new(big.Int) //.Mul(curve.Params().N, new(big.Int).SetInt64(int64(iter/2)))

	Rx.Add(Rx, sig_R)
	if Rx.Cmp(curve.Params().P) != -1 {
		return nil, errors.New("calculated Rx is larger than curve P")
	}

	// convert 02<Rx> to point R. (step 1.2 and 1.3). If we are on an odd
	// iteration then 1.6 will be done with -R, so we calculate the other
	// term when uncompressing the point.
	Ry, err := decompressPoint(Rx, iter%2 == 1)
	if err != nil {
		return nil, err
	}

	// 1.4 Check n*R is point at infinity
	if doChecks {
		nRx, nRy := curve.ScalarMult(Rx, Ry, curve.Params().N.Bytes())
		if nRx.Sign() != 0 || nRy.Sign() != 0 {
			return nil, errors.New("n*R does not equal the point at infinity")
		}
	}

	// 1.5 calculate e from message using the same algorithm as ecdsa
	// signature calculation.
	//e := hashToInt(msg)
	// sha256的hash刚好是32字节，256bit
	// 因此可以直接转换为 e
	e := new(big.Int).SetBytes(msg)

	// Step 1.6.1:
	// We calculate the two terms sR and eG separately multiplied by the
	// inverse of r (from the signature). We then add them to calculate
	// Q = r^-1(sR-eG)
	invr := new(big.Int).ModInverse(sig_R, curve.Params().N)

	// first term.
	invrS := new(big.Int).Mul(invr, sig_S)
	invrS.Mod(invrS, curve.Params().N)
	sRx, sRy := curve.ScalarMult(Rx, Ry, invrS.Bytes())

	// second term.
	e.Neg(e)
	e.Mod(e, curve.Params().N)
	e.Mul(e, invr)
	e.Mod(e, curve.Params().N)
	minuseGx, minuseGy := curve.ScalarBaseMult(e.Bytes())

	// TODO(oga) this would be faster if we did a mult and add in one
	// step to prevent the jacobian conversion back and forth.
	Qx, Qy := curve.Add(sRx, sRy, minuseGx, minuseGy)

	return &ecdsa.PublicKey{
		Curve: curve,
		X:     Qx,
		Y:     Qy,
	}, nil
}

// decompressPoint decompresses a point on the given curve given the X point and
// the solution to use.
func decompressPoint(x *big.Int, ybit bool) (*big.Int, error) {
	// TODO(oga) This will probably only work for secp256k1 due to
	// optimizations.

	curve := secp256k1.S256()
	// Y = +-sqrt(x^3 + B)
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)
	x3.Add(x3, curve.Params().B)

	// now calculate sqrt mod p of x2 + B
	// This code used to do a full sqrt based on tonelli/shanks,
	// but this was replaced by the algorithms referenced in
	// https://bitcointalk.org/index.php?topic=162805.msg1712294#msg1712294
	y := new(big.Int).Exp(x3, curve.QPlus1Div4(), curve.Params().P)

	// y坐标是否需要根据X轴翻转
	if ybit != isOdd(y) {
		y.Sub(curve.Params().P, y)
	}
	if ybit != isOdd(y) {
		return nil, fmt.Errorf("ybit doesn't match oddness")
	}
	return y, nil
}

// 是否为奇数
func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}
