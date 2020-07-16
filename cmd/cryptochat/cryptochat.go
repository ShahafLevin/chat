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
	"log"

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
		log.Printf("Failed to generate private/public key pair: %s\n", err)
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
func EncryptMessage(key Key, plaintext []byte) (cipherblob []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherblock := NewBlock(plaintext)

	mode := cipher.NewCBCEncrypter(block, cipherblock.Iv)
	mode.CryptBlocks(cipherblock.Ciphertext, cipherblock.Ciphertext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	if cipherblob, err = Marshal((*cipherblock)); err != nil {
		panic(err)
	}

	return cipherblob
}

// DecryptMessage deecrypts a messgae using a given key
func DecryptMessage(key Key, cipherblob []byte) (plaintext []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	var cipherblock Block
	if cipherblock, err = UnMarshal(cipherblob); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, cipherblock.Iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(cipherblock.Ciphertext, cipherblock.Ciphertext)

	return cipherblock.Ciphertext[:cipherblock.PaddingLen]
}
