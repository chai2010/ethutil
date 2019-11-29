// 以太坊工具箱(零第三方库依赖) 版权 @2019 柴树杉。

package ethutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"

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

	// AES加密数据
	var aesCTRXOR = func(key, inText, iv []byte) ([]byte, error) {
		aesBlock, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		stream := cipher.NewCTR(aesBlock, iv)
		outText := make([]byte, len(inText))
		stream.XORKeyStream(outText, inText)
		return outText, err
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
	return "", "", fmt.Errorf("todo")
}

func (p *ksKeyJSON) decryptKeyV3(password string) (uuid, privateKey string, err error) {
	return "", "", fmt.Errorf("todo")
}

/*
func decryptKeyV3(keyProtected *encryptedKeyJSONV3, auth string) (keyBytes []byte, keyId []byte, err error) {
	if keyProtected.Version != version {
		return nil, nil, fmt.Errorf("version not supported: %v", keyProtected.Version)
	}
	keyId = uuid.Parse(keyProtected.Id)
	plainText, err := DecryptDataV3(keyProtected.Crypto, auth)
	if err != nil {
		return nil, nil, err
	}
	return plainText, keyId, err
}

func DecryptDataV3(cryptoJson CryptoJSON, auth string) ([]byte, error) {
	if cryptoJson.Cipher != "aes-128-ctr" {
		return nil, fmt.Errorf("cipher not supported: %v", cryptoJson.Cipher)
	}
	mac, err := hex.DecodeString(cryptoJson.MAC)
	if err != nil {
		return nil, err
	}

	iv, err := hex.DecodeString(cryptoJson.CipherParams.IV)
	if err != nil {
		return nil, err
	}

	cipherText, err := hex.DecodeString(cryptoJson.CipherText)
	if err != nil {
		return nil, err
	}

	derivedKey, err := getKDFKey(cryptoJson, auth)
	if err != nil {
		return nil, err
	}

	calculatedMAC := crypto.Keccak256(derivedKey[16:32], cipherText)
	if !bytes.Equal(calculatedMAC, mac) {
		return nil, ErrDecrypt
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		return nil, err
	}
	return plainText, err
}

*/
