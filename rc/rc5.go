package rc

import (
	"crypto/rand"
	"fmt"

	"github.com/dgryski/go-rc5"
)

func Rc5() {
	key := make([]byte, 16) // 16 byte rc4 key - variable length
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	cipherText := encryptrc5([]byte("This is a secret message from mustafa!!"), key) // Encrypt
	fmt.Println(string(decryptrc5(cipherText, key)))                                 // Decrypt
}

func encryptrc5(message, key []byte) []byte {
	cipherText := make([]byte, len(message))
	rc, err := rc5.New(key)
	if err != nil {
		panic(err)
	}

	rc.Encrypt(cipherText, message)
	return cipherText
}

func decryptrc5(message, key []byte) []byte {
	cipherText := make([]byte, len(message))
	rc, err := rc5.New(key)
	if err != nil {
		panic(err)
	}

	rc.Decrypt(cipherText, message)
	return cipherText
}
