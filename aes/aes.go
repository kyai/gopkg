package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
const Key = "6666666666666666"

func EnToken(s string) string {
	result, err := AesEncrypt([]byte(s), []byte(Key))
	if err != nil {
		// panic(err)
		return ""
	}

	ss := base64.StdEncoding.EncodeToString(result)

	ss = strings.Replace(ss, "/", "_a", -1)
	ss = strings.Replace(ss, "+", "_b", -1)
	ss = strings.Replace(ss, "=", "_c", -1)

	return ss
}

func DeToken(s string) string {
	s = strings.Replace(s, "_a", "/", -1)
	s = strings.Replace(s, "_b", "+", -1)
	s = strings.Replace(s, "_c", "=", -1)

	ss, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}

	result, err := AesDecrypt(ss, []byte(Key))
	if err != nil {
		// panic(err)
		return ""
	}

	return string(result)
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ParseToArray(s string) map[string]interface{} {
	arr := strings.Split(s, "&")
	rdata := make(map[string]interface{})
	if len(s) < 2 {
		return rdata
	}
	for _, value := range arr {
		v_arr := strings.Split(value, "=")
		rdata[v_arr[0]] = v_arr[1]
	}
	return rdata
}

func ParseToArrayJson(s string) map[string]interface{} {
	rdata := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &rdata)
	if err != nil {
		return nil
	}
	return rdata
}
