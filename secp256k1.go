// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import "math/big"

// secp256k1 椭圆曲线参数
const (
	// secp256k1 模素数的不同形式
	// SECP256K1_P0 是指数组成的常量表达式(无法在运行时直接使用)
	// SECP256K1_P1 是最终的十进制格式(无法在运行时直接使用)
	// SECP256K1_P2 是十进制的字符串面值
	// SECP256K1_P0/SECP256K1_P1 和 SECP256K1_P3 无法在运行时直接比较
	//
	// var P, _ = new(big.Int).SetString(SECP256K1_P2, 10)
	// assert(t, P.String() == SECP256K1_P2)
	//
	SECP256K1_P0 = 1<<256 - 1<<32 - 1<<9 - 1<<8 - 1<<7 - 1<<6 - 1<<4 - 1
	SECP256K1_P1 = 115792089237316195423570985008687907853269984665640564039457584007908834671663
	SECP256K1_P2 = "115792089237316195423570985008687907853269984665640564039457584007908834671663"
)

// secp256k1 模素数(常量)
var (
	_SECP256K1_P, _ = new(big.Int).SetString(SECP256K1_P2, 10)
)

// 根据 x 计算 secp256k1 椭圆曲线函数的值
// secp256k1_Fx(x) = (x^3 + 7)% P
func Secp256k1_Fx(x *big.Int) *big.Int {
	v := new(big.Int)
	v.Exp(x, big.NewInt(3), nil).Add(v, big.NewInt(7)).Mod(v, _SECP256K1_P)
	return v
}

// 根据 y 计算 secp256k1 椭圆曲线函数的值
// secp256k1_Fy(y) = (y^2)% P
func Secp256k1_Fy(y *big.Int) *big.Int {
	v := new(big.Int)
	v.Exp(y, big.NewInt(2), nil).Mod(v, _SECP256K1_P)
	return v
}

// secp256k1_Fx(x) - secp256k1_Fy(y)
// 如果结果为0, 表示(x,y)点在椭圆曲线上(z.BitLen() == 0)
func Secp256k1_Fx_Fy(x, y *big.Int) *big.Int {
	v := new(big.Int)
	v.Sub(Secp256k1_Fx(x), Secp256k1_Fy(y)).Mod(v, _SECP256K1_P)
	return v
}
