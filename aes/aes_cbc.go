package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AES_CBC() {
	key := generateAESKey(32)
	fmt.Println("key(bytes): ", key)
	fmt.Println("key(string): ", string(key))

	iv := generateIV(aes.BlockSize)
	fmt.Println("iv(bytes): ", iv)
	fmt.Println("iv(string): ", string(iv))

	plainText := "Hello this is mustafa!!!! and I am testing aes encryption using ecb mode \nHello this is mustafa!!! and I am testing aes encryption using ecb mode \nI repeated same sentence twice just to see the similarity of ciphertext"
	cipher := aes_cbc_encrypt([]byte(plainText), key, iv)
	fmt.Println("cipherText(bytes): ", cipher)
	fmt.Println("cipherText(string): ", string(cipher))

	plain := aes_cbc_decrypt(cipher, key, iv)
	fmt.Println("decryptedText(bytes): ", plain)
	fmt.Println("decryptedText(string): ", string(plain))
}

func aes_cbc_encrypt(data, key, iv []byte) []byte {
	blck, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	data = pkcs7Pad(data, aes.BlockSize)
	mode := cipher.NewCBCEncrypter(blck, iv)
	cipherText := make([]byte, len(data))
	mode.CryptBlocks(cipherText, data)

	return cipherText
}

func aes_cbc_decrypt(data, key, iv []byte) []byte {
	blck, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(blck, iv)
	plainText := make([]byte, len(data))
	mode.CryptBlocks(plainText, data)

	return pkcs7Unpad(plainText, aes.BlockSize)
}
