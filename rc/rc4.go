package rc

import (
	"crypto/rand"
	"crypto/rc4"
	"fmt"
)

func Rc4() {
	key := make([]byte, 32) // 32 byte rc4 key - variable length
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	cipherText := encryptrc4([]byte("This is a secret message from mustafa!!"), key) // Encrypt
	fmt.Println(string(encryptrc4(cipherText, key)))                                 // Decrypt
}

func encryptrc4(message, key []byte) []byte {
	cipherText := make([]byte, len(message))
	rc, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	rc.XORKeyStream(cipherText, message)
	return cipherText
}
