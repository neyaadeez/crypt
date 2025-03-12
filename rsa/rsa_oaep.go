package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func RSA_OAEP() {
	//key gen
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("keygen err: ", err.Error())
		return
	}

	//encryption
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, []byte("Hello this is secret message"), nil)
	if err != nil {
		fmt.Println("encryption err: ", err)
		return
	}
	fmt.Println("CipherText: ", string(ciphertext))

	//decryption
	originalMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Println("decription err: ", err)
		return
	}
	fmt.Println("originalMessage: ", string(originalMessage))
}
