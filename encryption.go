package solidgate

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func pkcs7Pad(b []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("block size less than 0")
	}
	if b == nil || len(b) == 0 {
		return nil, errors.New("empty data to encrypt")
	}
	n := blockSize - (len(b) % blockSize)
	pb := make([]byte, len(b)+n)

	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))

	return pb, nil
}

func EncryptCBC(key, data, iv []byte) ([]byte, error) {
	data, err := pkcs7Pad(data, aes.BlockSize)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[:len(data)], data)

	return ciphertext[:len(data)], nil
}
func DecryptCBC(key, data, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	return data, nil
}
