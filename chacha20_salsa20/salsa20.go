package chacha20salsa20

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/salsa20"
)

func Salsa20() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	nounce := make([]byte, 8)
	_, err = rand.Read(nounce)
	if err != nil {
		panic(err)
	}

	cipherText := encryptSala20([]byte("This is a secret message from mustafa!!"), key, nounce) // Encrypt
	fmt.Println(string(decryptSalsa20(cipherText, key, nounce)))                                // Decrypt
}

func encryptSala20(message, key, nounce []byte) []byte {
	cipherText := make([]byte, len(message))

	salsa20.XORKeyStream(cipherText, message, nounce, (*[32]byte)(key))
	return cipherText
}

func decryptSalsa20(cipherText, key, nounce []byte) []byte {
	plaintext := make([]byte, len(cipherText))

	salsa20.XORKeyStream(plaintext, cipherText, nounce, (*[32]byte)(key))
	return plaintext
}
