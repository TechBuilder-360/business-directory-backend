package unit

import (
	"encoding/hex"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"testing"
)

func TestEncryption(t *testing.T) {
	randomstring := utility.GenerateRandomString(32)
	key := hex.EncodeToString(randomstring)
	fmt.Printf("Key: %s\n", key)
	plaintext := "Test string"

	// Encrypt given text, with key
	ciphertext, err := utility.Encrypt(key, plaintext)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Cipher Text: %s\n", ciphertext)

	// Decrypt cipher text with key
	text, err := utility.Decrypt(key, ciphertext)
	if err != nil {
		t.Error(err.Error())
	}

	if text != plaintext {
		t.Errorf("AES encryption test failed. Returned %s != %s", plaintext, text)
	}
}