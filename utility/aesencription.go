package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/zenazn/pkcs7pad"
	"io"
	random "math/rand"
	"strings"
	"time"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-)([]}{.?:><")
var numeric = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func GenerateRandomString(size uint) []byte {
	random.Seed(time.Now().UnixNano())

	b := make([]byte, size)
	for i := range b {
		b[i] = letters[random.Intn(len(letters))]
	}

	return b
}

func GenerateNumericToken(size int) string {
	random.Seed(time.Now().UnixNano())

	b := make([]string, size)
	for i := range b {
		b[i] = numeric[random.Intn(len(numeric))]
	}

	return strings.Join(b, "")
}

func Encrypt(key string, text string) (string, error) {
	keyByte, _ := hex.DecodeString(key)
	plaintext := []byte(text)

	// Pad plain text if length isn't of block size 16
	plaintext = pkcs7pad.Pad(plaintext, aes.BlockSize)

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, it would be appended to the head of cipher text
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(key string, text string) (string, error) {
	keyByte, _ := hex.DecodeString(key)
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	// The IV extracted from the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext, err := pkcs7pad.Unpad(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
