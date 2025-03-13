package aes

import (
	"crypto/aes"
	"fmt"
)

func AES_ECB() {
	key := generateAESKey(32)
	fmt.Println("key(bytes): ", key)
	fmt.Println("key(string): ", string(key))

	plainText := "Hello this is mustafa!!!! and I am testing aes encryption using ecb mode \nHello this is mustafa!!! and I am testing aes encryption using ecb mode \nI repeated same sentence twice just to see the similarity of ciphertext"
	cipher := aes_ecb_encrypt([]byte(plainText), key)
	fmt.Println("cipherText(bytes): ", cipher)
	fmt.Println("cipherText(string): ", string(cipher))

	plain := aes_ecb_decrypt(cipher, key)
	fmt.Println("decryptedText(bytes): ", plain)
	fmt.Println("decryptedText(string): ", string(plain))
}

func aes_ecb_encrypt(data []byte, key []byte) []byte {
	blck, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	paddedData := pkcs7Pad(data, aes.BlockSize)
	cipherText := make([]byte, len(paddedData))
	for i := 0; i < len(data); i += aes.BlockSize {
		blck.Encrypt(cipherText[i:i+aes.BlockSize], data[i:i+aes.BlockSize])
	}
	return cipherText
}

func aes_ecb_decrypt(cipherText []byte, key []byte) []byte {
	blck, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plainText := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i += aes.BlockSize {
		blck.Decrypt(plainText[i:i+aes.BlockSize], cipherText[i:i+aes.BlockSize])
	}

	unPaddedData := pkcs7Unpad(plainText, aes.BlockSize)

	return unPaddedData
}
