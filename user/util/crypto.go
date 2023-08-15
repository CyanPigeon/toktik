package util

import (
	"crypto/aes"
	cipher "crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type cryptoConfig struct {
	Key []byte
}

var cfg = cryptoConfig{
	Key: []byte("CyanPigeon"),
}

func AESEncrypt(s string) (string, error) {
	block, err := aes.NewCipher(cfg.Key)
	if err != nil {
		return "", err
	}
	bs := []byte(s)
	cipherText := make([]byte, aes.BlockSize+len(bs))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], bs)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AESDecrypt(s string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(cfg.Key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
