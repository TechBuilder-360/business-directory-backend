package utility

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"io/ioutil"
	"net/mail"
	"strings"
	"time"
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

// TitleCase ...
func TitleCase(text string) string {
	return strings.Title(text)
}

func FormatDate(date time.Time) string {
	return date.Format("02-01-2006")
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CapitalizeFirstCharacter(s string) string {
	return cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.ToLower(strings.TrimSpace(s)))
}

func GenerateUUID() string {
	return uuid.NewString()
}
