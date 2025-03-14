package ecc

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func ECC_DS() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("publicKey size: ", len(publicKey))
	fmt.Println("privateKey size: ", len(privateKey))

	message := "Hello this is Mustafa!"
	signature := ed25519.Sign(privateKey, []byte(message))
	fmt.Println("signature(hex): ", hex.EncodeToString(signature))
	fmt.Println("signature(base64): ", base64.RawStdEncoding.EncodeToString(signature))

	fmt.Println("signature size: ", len(signature))

	valid := ed25519.Verify(publicKey, []byte(message), signature)
	if valid {
		fmt.Println("signature is valid!!!")
	} else {
		fmt.Println("signature is Invalid!!!")
	}

}
