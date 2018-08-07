package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// length(bytes)	type
// 16				128
// 24				192
// 32				256
var key []byte

func SetKey(s string) {
	key = []byte(s)
}

func Encrypt(s string) (result string, err error) {
	var ss []byte
	ss, err = AesEncrypt([]byte(s), key)
	if err != nil {
		return
	}
	result = base64.StdEncoding.EncodeToString(ss)
	return
}

func Decrypt(s string) (result string, err error) {
	var ss []byte
	if ss, err = base64.StdEncoding.DecodeString(s); err != nil {
		return
	}
	if ss, err = AesDecrypt(ss, key); err != nil {
		return
	}
	result = string(ss)
	return
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
