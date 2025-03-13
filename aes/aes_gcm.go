package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AES_GCM() {
	key := generateAESKey(32)
	fmt.Println("key(bytes): ", key)
	fmt.Println("key(string): ", string(key))

	nounce := generateIV(12) // AES-GCM requires a nonce of at least 12 bytes (recommended by NIST).
	fmt.Println("iv(bytes): ", nounce)
	fmt.Println("iv(string): ", string(nounce))

	plainText := "Hello this is mustafa!!!! and I am testing aes encryption using ecb mode \nHello this is mustafa!!! and I am testing aes encryption using ecb mode \nI repeated same sentence twice just to see the similarity of ciphertext"
	cipher := aes_gcm_encrypt([]byte(plainText), key, nounce)
	fmt.Println("cipherText(bytes): ", cipher)
	fmt.Println("cipherText(string): ", string(cipher))

	plain := aes_gcm_decrypt(cipher, key, nounce)
	fmt.Println("decryptedText(bytes): ", plain)
	fmt.Println("decryptedText(string): ", string(plain))
}

func aes_gcm_encrypt(plaintext, key, nonce []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	return ciphertext
}

func aes_gcm_decrypt(ciphertext, key, nonce []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic("Decryption failed!")
	}
	return plaintext
}
