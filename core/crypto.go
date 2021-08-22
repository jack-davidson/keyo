package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func Seed(size int32) []byte {
	randSeed := make([]byte, size)
	rand.Read(randSeed)
	seed := sha512.Sum512(randSeed)
	return seed[:]
}

func Derive(complement []byte, seed []byte) []byte {
	passphrase := sha512.Sum512(seed)
	return passphrase[:]
}

func Encode(passphrase []byte) string {
	return base64.StdEncoding.EncodeToString(passphrase)
}
