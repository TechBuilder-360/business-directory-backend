package utility

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

// ExtractRequestBody returns body and body reader
func ExtractRequestBody(c io.ReadCloser) (string, io.ReadCloser) {
	// Read the Body content
	var bodyBytes []byte
	if c != nil {
		bodyBytes, _ = ioutil.ReadAll(c)
	}
	return string(bodyBytes), ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

// ComputeHmac256 ...
func ComputeHmac256(message string, secret string) string {
	key, _ := hex.DecodeString(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// UserHasRole ...
func UserHasRole(userRole, requiredRole []string) bool {
	for _, v := range requiredRole {
		for _, b := range userRole {
			if v == b {
				return true
			}
		}
	}
	return false
}

func StringPtrToString(str *string) string {
	if str != nil {
		return *str
	}

	return ""
}
