package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var KEY = "xiaodipy20181688"

func AesEncrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(KEY))
	if nil != err {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, []byte(KEY)[:blockSize])
	origData = pkcs5Padding(origData, blockSize)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(KEY))
	if nil != err {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(KEY)[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Base64AesDecrypt(data string) string {
	dst, _ := base64.StdEncoding.DecodeString(data)
	dst, _ = AesDecrypt(dst)
	return string(dst)
}
