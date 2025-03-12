package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func RAS_PSS() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("keyGen error: ", err.Error())
		return
	}

	hashed := sha256.Sum256([]byte("this is a test!!!"))
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		fmt.Println("signing error: ", err.Error())
		return
	}
	fmt.Println(string(signature))

	receivedMessageHash := sha256.Sum256([]byte("this is a test!!!"))
	err = rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA256, receivedMessageHash[:], signature, nil)
	if err != nil {
		fmt.Println("signature is not valid or the message is not signed by the user: ", err.Error())
		return
	}
	fmt.Println("signature and message is verified!!!")
}
