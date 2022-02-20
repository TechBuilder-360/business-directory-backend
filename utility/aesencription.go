package utility

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"

	"github.com/TechBuilder-360/business-directory-backend.git/configs"
)


func Encrypt(body []byte) string {

	Config := configs.Configuration()
	
        // create cipher
    c, err := aes.NewCipher([]byte(Config.AesKey))
    CheckError(err)
        
        // allocate space for ciphered data
    out := make([]byte, len(body))
 
        // encrypt
    c.Encrypt(out, []byte(body))
        // return hex string
    return hex.EncodeToString(out)
}

func Decrypt(ct string) {
	Config := configs.Configuration()
	ciphertext, _ := hex.DecodeString(ct)
     
	c, err := aes.NewCipher([]byte(Config.AesKey))
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