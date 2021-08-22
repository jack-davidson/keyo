package core

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	seed := Seed(1024)
	fmt.Println(hex.EncodeToString(seed))
	passphrase := Derive([]byte("hello"), seed)
	fmt.Println(hex.EncodeToString(passphrase))
	encodedPassphrase := Encode(passphrase)
	fmt.Println(encodedPassphrase)
}
