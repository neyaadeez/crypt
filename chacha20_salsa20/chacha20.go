package chacha20salsa20

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20"
)

func ChaCha20() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	nounce := make([]byte, 12)
	_, err = rand.Read(nounce)
	if err != nil {
		panic(err)
	}

	cipherText := encryptChaCha20([]byte("This is a secret message from mustafa!!"), key, nounce) // Encrypt
	fmt.Println(string(decryptChaCha20(cipherText, key, nounce))) // Decrypt
}

func encryptChaCha20(message, key, nounce []byte) []byte {
	cipherText := make([]byte, len(message))
	chaChaCipher, err := chacha20.NewUnauthenticatedCipher(key, nounce)
	if err != nil {
		panic(err)
	}

	chaChaCipher.XORKeyStream(cipherText, message)
	return cipherText
}

func decryptChaCha20(cipherText, key, nounce []byte) []byte {
	plaintext := make([]byte, len(cipherText))
	chaChaCipher, err := chacha20.NewUnauthenticatedCipher(key, nounce)
	if err != nil {
		panic(err)
	}

	chaChaCipher.XORKeyStream(plaintext, cipherText)
	return plaintext
}
