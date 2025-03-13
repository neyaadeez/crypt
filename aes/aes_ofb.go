package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AES_OFB() {
	key := generateAESKey(32)
	fmt.Println("key(bytes): ", key)
	fmt.Println("key(string): ", string(key))

	iv := generateIV(aes.BlockSize)
	fmt.Println("iv(bytes): ", iv)
	fmt.Println("iv(string): ", string(iv))

	plainText := "Hello this is mustafa!!!! and I am testing aes encryption using ecb mode \nHello this is mustafa!!! and I am testing aes encryption using ecb mode \nI repeated same sentence twice just to see the similarity of ciphertext"
	cipher := aes_ofb_encrypt([]byte(plainText), key, iv)
	fmt.Println("cipherText(bytes): ", cipher)
	fmt.Println("cipherText(string): ", string(cipher))

	plain := aes_ofb_decrypt(cipher, key, iv)
	fmt.Println("decryptedText(bytes): ", plain)
	fmt.Println("decryptedText(string): ", string(plain))
}

func aes_ofb_encrypt(plaintext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func aes_ofb_decrypt(ciphertext, key, iv []byte) []byte {
	return aes_ofb_encrypt(ciphertext, key, iv) // OFB uses same function for encryption/decryption
}
