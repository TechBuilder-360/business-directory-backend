package utility

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func ExtractRequestBody(c *gin.Context) (string, io.ReadCloser) {
	// Read the Body content
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	return string(bodyBytes), ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}


func ComputeHmac256(message string, secret string) string {
	key, _ := base64.StdEncoding.DecodeString(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}