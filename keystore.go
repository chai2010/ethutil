// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

// KeyStore对应的JSON结构
// V1和V3基本是同样的结构
type ksKeyJSON struct {
	Address string       `json:"address"`
	Crypto  ksCryptoJSON `json:"crypto"`
	Id      string       `json:"id"`
	Version int          `json:"version"`
}

// 加密部分的参数
type ksCryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams ksCipherparamsJSON     `json:"cipherparams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfparams"`
	MAC          string                 `json:"mac"`
}

// 私钥加密算法需要的参数
type ksCipherparamsJSON struct {
	IV string `json:"iv"`
}

// 加密私钥
// 生成KeyStore格式JSON
func KeyStoreEncryptKey(uuid, privateKey, password string) (keyjson []byte, err error) {
	// 生成key的长度
	const scryptDKLen = 32

	// 默认的生成key参数, 可以调整
	const scryptN = 262144
	const scryptR = 8
	const scryptP = 1

	// 采用 scrypt 算法从 password 生成解密 key
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	derivedKey, err := scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return nil, err
	}

	// 取生成密码的后16字节作为解密Key
	encryptKey := derivedKey[:16]

	// 生成AES解密参数
	iv := make([]byte, aes.BlockSize) // 16
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}

	// 加密私钥
	keyBytes := AsBigint(privateKey, 64).Bytes()
	cipherText, err := aesCTRXOR(encryptKey, keyBytes, iv)
	if err != nil {
		return nil, err
	}

	// 记录加密私钥的参数
	cryptoStruct := ksCryptoJSON{
		// 私钥加密算法
		Cipher: "aes-128-ctr",
		// 私钥加密后的数据
		CipherText: fmt.Sprintf("%x", cipherText),
		// 私钥加密算法参数
		CipherParams: ksCipherparamsJSON{
			IV: fmt.Sprintf("%x", iv),
		},

		// 私钥加密算法key的生成算法
		KDF: "scrypt",
		// 生成key算法的参数
		KDFParams: map[string]interface{}{
			"n":     scryptN,
			"r":     scryptR,
			"p":     scryptP,
			"dklen": scryptDKLen,
			"salt":  fmt.Sprintf("%x", salt),
		},

		// password 校验码
		MAC: Keccak256Hash(derivedKey[16:32], cipherText),
	}

	// 生成 JSON
	return json.Marshal(&ksKeyJSON{
		Address: GenAddressFromPrivateKey(privateKey),
		Crypto:  cryptoStruct,
		Id:      uuid,
		Version: 3,
	})
}

// 解密私钥
func KeyStoreDecryptKey(keyjson []byte, password string) (uuid, privateKey string, err error) {
	// 解码JSON
	var k ksKeyJSON
	if err := json.Unmarshal(keyjson, &k); err != nil {
		return "", "", err
	}

	// 处理不同版本
	switch {
	case k.Version == 1:
		return k.decryptKeyV1(password)
	case k.Version == 3:
		return k.decryptKeyV3(password)
	default:
		return k.decryptKeyV3(password)
	}
}

func (p *ksKeyJSON) decryptKeyV1(password string) (uuid, privateKey string, err error) {
	if p.Crypto.Cipher != "aes-128-ctr" {
		return "", "", fmt.Errorf("cipher not supported: %v", p.Crypto.Cipher)
	}

	// 重现加密key
	derivedKey, err := p.getKDFKey(password)
	if err != nil {
		return "", "", err
	}

	// 校验 password 是否正确
	calculatedMAC := Keccak256Hash(derivedKey[16:32], []byte(p.Crypto.CipherText))
	if calculatedMAC != p.Crypto.MAC {
		return "", "", errors.New("could not decrypt key with given password")
	}

	// AES 对称解密
	plainText, err := aesCBCDecrypt(
		derivedKey[:16],
		[]byte(p.Crypto.CipherText),
		[]byte(p.Crypto.CipherParams.IV),
	)
	if err != nil {
		return "", "", err
	}

	// OK
	return p.Id, fmt.Sprintf("%x", plainText), err
}

func (p *ksKeyJSON) decryptKeyV3(password string) (uuid, privateKey string, err error) {
	if p.Crypto.Cipher != "aes-128-ctr" {
		return "", "", fmt.Errorf("cipher not supported: %v", p.Crypto.Cipher)
	}

	// 重现加密key
	derivedKey, err := p.getKDFKey(password)
	if err != nil {
		return "", "", err
	}

	// 校验 password 是否正确
	calculatedMAC := Keccak256Hash(derivedKey[16:32], []byte(p.Crypto.CipherText))
	if calculatedMAC != p.Crypto.MAC {
		return "", "", errors.New("could not decrypt key with given password")
	}

	// AES 对称解密
	plainText, err := aesCTRXOR(
		derivedKey[:16],
		[]byte(p.Crypto.CipherText),
		[]byte(p.Crypto.CipherParams.IV),
	)
	if err != nil {
		return "", "", err
	}

	// OK
	return p.Id, fmt.Sprintf("%x", plainText), err
}

func (p *ksKeyJSON) getKDFKey(password string) ([]byte, error) {
	authArray := []byte(password)
	salt, err := hex.DecodeString(p.Crypto.KDFParams["salt"].(string))
	if err != nil {
		return nil, err
	}

	switch {
	case p.Crypto.KDF == "scrypt":
		return scrypt.Key(authArray, salt,
			p.getKDFParamsInt("n"),
			p.getKDFParamsInt("r"),
			p.getKDFParamsInt("p"),
			p.getKDFParamsInt("dklen"),
		)

	case p.Crypto.KDF == "pbkdf2":
		if prf := p.getKDFParamsString("prf"); prf != "hmac-sha256" {
			return nil, fmt.Errorf("unsupported PBKDF2 PRF: %s", prf)
		}
		key := pbkdf2.Key(authArray, salt,
			p.getKDFParamsInt("c"),
			p.getKDFParamsInt("dklen"),
			sha256.New,
		)
		return key, nil

	default:
		return nil, fmt.Errorf("unsupported KDF: %s", p.Crypto.KDF)
	}
}

func (p *ksKeyJSON) getKDFParamsInt(name string) int {
	f, _ := strconv.ParseFloat(p.getKDFParamsString(name), 64)
	return int(f)
}
func (p *ksKeyJSON) getKDFParamsString(name string) string {
	return fmt.Sprint(p.Crypto.KDFParams[name])
}

// AES加密/解密数据
func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesBlock, iv)
	outText := make([]byte, len(inText))
	stream.XORKeyStream(outText, inText)
	return outText, err
}

// V1版本加密/解密
func aesCBCDecrypt(key, cipherText, iv []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypter := cipher.NewCBCDecrypter(aesBlock, iv)
	paddedPlaintext := make([]byte, len(cipherText))
	decrypter.CryptBlocks(paddedPlaintext, cipherText)
	plaintext := pkcs7Unpad(paddedPlaintext)
	if plaintext == nil {
		return nil, errors.New("could not decrypt key with given password")
	}
	return plaintext, err
}

// From https://leanpub.com/gocrypto/read#leanpub-auto-block-cipher-modes
func pkcs7Unpad(in []byte) []byte {
	if len(in) == 0 {
		return nil
	}

	padding := in[len(in)-1]
	if int(padding) > len(in) || padding > aes.BlockSize {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
		if in[i] != padding {
			return nil
		}
	}
	return in[:len(in)-int(padding)]
}
