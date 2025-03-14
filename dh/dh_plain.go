package dh

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Generate a private key (random number)
func generatePrivateKey(prime *big.Int) *big.Int {
	privateKey, err := rand.Int(rand.Reader, prime)
	if err != nil {
		panic(err)
	}
	return privateKey
}

// Compute public key: g^privateKey mod p
func computePublicKey(base, privateKey, prime *big.Int) *big.Int {
	publicKey := new(big.Int).Exp(base, privateKey, prime)
	return publicKey
}

// Compute shared secret: otherPublicKey^privateKey mod p
func computeSharedSecret(otherPublicKey, privateKey, prime *big.Int) *big.Int {
	sharedSecret := new(big.Int).Exp(otherPublicKey, privateKey, prime)
	return sharedSecret
}

func DH_Plain() {
	prime, _ := new(big.Int).SetString("23", 10) // Prime number (p)
	base := big.NewInt(5)                        // Base (g)

	// Alice generates private & public key
	alicePrivate := generatePrivateKey(prime)
	alicePublic := computePublicKey(base, alicePrivate, prime)

	// Bob generates private & public key
	bobPrivate := generatePrivateKey(prime)
	bobPublic := computePublicKey(base, bobPrivate, prime)

	// Exchange public keys & compute shared secret
	aliceShared := computeSharedSecret(bobPublic, alicePrivate, prime)
	bobShared := computeSharedSecret(alicePublic, bobPrivate, prime)

	// Print results
	fmt.Println("Alice Private Key:", alicePrivate)
	fmt.Println("Alice Public Key:", alicePublic)
	fmt.Println("Bob Private Key:", bobPrivate)
	fmt.Println("Bob Public Key:", bobPublic)
	fmt.Println("Alice's Shared Secret:", aliceShared)
	fmt.Println("Bob's Shared Secret:  ", bobShared)

	// Verify both parties have the same shared secret
	if aliceShared.Cmp(bobShared) == 0 {
		fmt.Println("✅ Diffie-Hellman key exchange successful!")
	} else {
		fmt.Println("❌ Key exchange failed!")
	}
}
