package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func RSA_PKCS() {
	//key gen
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("keygen err: ", err.Error())
		return
	}

	//encryption
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, []byte("Hello this is secret message"))
	if err != nil {
		fmt.Println("encryption err: ", err)
		return
	}
	fmt.Println("CipherText: ", string(ciphertext))

	//decryption
	originalMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		fmt.Println("decription err: ", err)
		return
	}
	fmt.Println("originalMessage: ", string(originalMessage))

}
