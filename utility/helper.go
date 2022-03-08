package utility

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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

func ComputeHmac256(message string, secret string) string {
	key, _ := base64.StdEncoding.DecodeString(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func IsContain(role []string, data ...string) bool {

	for _, v := range role {
		for _, b := range data {
			if v == b {
				return true
			}
		}

	}

	return false
}
