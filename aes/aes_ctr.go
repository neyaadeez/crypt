package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AES_CTR() {
	key := generateAESKey(32)
	fmt.Println("key(bytes): ", key)
	fmt.Println("key(string): ", string(key))

	nounce := generateIV(aes.BlockSize)
	fmt.Println("iv(bytes): ", nounce)
	fmt.Println("iv(string): ", string(nounce))

	plainText := "Hello this is mustafa!!!! and I am testing aes encryption using ecb mode \nHello this is mustafa!!! and I am testing aes encryption using ecb mode \nI repeated same sentence twice just to see the similarity of ciphertext"
	cipher := aes_ctr_encrypt([]byte(plainText), key, nounce)
	fmt.Println("cipherText(bytes): ", cipher)
	fmt.Println("cipherText(string): ", string(cipher))

	plain := aes_ctr_decrypt(cipher, key, nounce)
	fmt.Println("decryptedText(bytes): ", plain)
	fmt.Println("decryptedText(string): ", string(plain))
}

func aes_ctr_encrypt(plaintext, key, nonce []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func aes_ctr_decrypt(ciphertext, key, nonce []byte) []byte {
	return aes_ctr_encrypt(ciphertext, key, nonce) // CTR mode is symmetric
}
