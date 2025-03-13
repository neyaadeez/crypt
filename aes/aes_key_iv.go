package aes

import (
	"bytes"
	"crypto/rand"
	"io"
)

func generateAESKey(size int) []byte {
	key := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		panic(err)
	}

	return key
}

func generateIV(size int) []byte {
	iv := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		panic(err)
	}

	return iv
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	paddedData := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(data, paddedData...)
}

func pkcs7Unpad(data []byte, blockSize int) []byte {
	lenght := len(data)
	padding := int(data[lenght-1])

	return data[:lenght-padding]
}
