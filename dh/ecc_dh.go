package dh

import (
	"bytes"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func ECC_DH() {
	// Alice
	alicePrivateKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	alicePublicKey := alicePrivateKey.PublicKey()
	fmt.Println("private key ALICE: ", hex.EncodeToString(alicePrivateKey.Bytes()))
	fmt.Println("public key ALICE: ", hex.EncodeToString(alicePublicKey.Bytes()))
	fmt.Println("private key ALICE(len): ", len(alicePrivateKey.Bytes()))
	fmt.Println("public key ALICE(len): ", len(alicePublicKey.Bytes()))

	// Bob
	bobPrivateKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	bobPublicKey := bobPrivateKey.PublicKey()
	fmt.Println("private key BOB: ", hex.EncodeToString(bobPrivateKey.Bytes()))
	fmt.Println("public key BOB: ", hex.EncodeToString(bobPublicKey.Bytes()))
	fmt.Println("private key BOB(len): ", len(bobPrivateKey.Bytes()))
	fmt.Println("public key BOB(len): ", len(bobPublicKey.Bytes()))

	// shared secret
	sharedSecretCalcByAlice, err := alicePrivateKey.ECDH(bobPublicKey)
	if err != nil {
		panic(err)
	}

	sharedSecretCalcByBob, err := bobPrivateKey.ECDH(alicePublicKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("shared secret by alice: ", base64.RawStdEncoding.EncodeToString(sharedSecretCalcByAlice))
	fmt.Println("shared secret by bob: ", base64.RawStdEncoding.EncodeToString(sharedSecretCalcByBob))
	fmt.Println("shared secret by alice(len): ", len(sharedSecretCalcByAlice))
	fmt.Println("shared secret by bob(len): ", len(sharedSecretCalcByBob))

	if bytes.Equal(sharedSecretCalcByAlice, sharedSecretCalcByBob) {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("NOT EQUAL")
	}
}
