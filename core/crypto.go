package core

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
	compHash := sha512.Sum512(complement)
	seedHash := sha512.Sum512(seed)
	passphrase := sha512.Sum512([]byte(string(compHash[:]) + string(seedHash[:])))
	return passphrase[:]
}

func Encode(passphrase []byte) string {
	return base64.StdEncoding.EncodeToString(passphrase)
}
