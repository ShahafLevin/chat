/*
Package cryptochat repsonsible to make the chat secure.
In order to do that, we use Elliptic Curve Diffie Helman (ECDH) for the key echange,
And AES to decrypt the messages themself. */
package cryptochat

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/aead/ecdh"
)

// Key is the simetric key we use to communicate with the user
type Key []byte

// ECDHKey contains the params for the ecdh key exchange
type ECDHKey struct {
	Curve       elliptic.Curve
	KeyExchange ecdh.KeyExchange
	Private     crypto.PrivateKey
	Public      crypto.PublicKey
}

// GenerateKey init the keys needed for the ecdh
func GenerateKey() (key *ECDHKey) {
	curve := elliptic.P256()
	p256 := ecdh.Generic(curve)

	private, public, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate private/public key pair: %s\n", err)
	}

	key = &ECDHKey{Curve: curve, KeyExchange: p256, Private: private, Public: public}
	return
}

// CreateAESKey creates the AES Key
func CreateAESKey(pub crypto.PublicKey, key ECDHKey) (secret []byte) {
	secret = key.KeyExchange.ComputeSecret(key.Private, pub)
	return
}

// EncryptMessage encrypt a messgae using a given key
func EncryptMessage(key Key, plaintext []byte) (ciphertext []byte) {
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	return ciphertext
}

// DecryptMessage deecrypts a messgae using a given key
func DecryptMessage(key Key, ciphertext []byte) (plaintext []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%s\n", ciphertext)
	return ciphertext
}
