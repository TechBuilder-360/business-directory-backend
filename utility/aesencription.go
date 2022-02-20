package utility

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)


func Encrypt(key []byte, plaintext string) string {
        // create cipher
    c, err := aes.NewCipher(key)
    CheckError(err)
        
        // allocate space for ciphered data
    out := make([]byte, len(plaintext))
 
        // encrypt
    c.Encrypt(out, []byte(plaintext))
        // return hex string
    return hex.EncodeToString(out)
}

func Decrypt(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)
     
	c, err := aes.NewCipher(key)
	CheckError(err)
     
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
     
	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
    }

func CheckError(err error) {
	if err != nil {
	    panic(err)
	}
    }