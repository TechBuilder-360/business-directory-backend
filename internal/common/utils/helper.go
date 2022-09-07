package utils

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/xid"

	"io"
	"io/ioutil"

	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func ToLower(text string) string {
	return strings.ToLower(text)
}

func CapitalizeFirstCharacter(s string) string {
	return cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.ToLower(strings.TrimSpace(s)))
}

func GenerateUUID() string {
	return uuid.NewString()
}

func GenerateUniqueID() string {
	guid := xid.New()

	return guid.String()
}

func ToMd5(text string) string {
	val := md5.Sum([]byte(text))

	return hex.EncodeToString(val[:])
}

func StructToMap(input interface{}, output interface{}) (interface{}, error) {

	err := mapstructure.Decode(input, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func AddToStr(add *string) string {
	if add == nil {
		return ""
	}

	return *add
}

func ImageUpload(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(configs.Instance.EnvCloudName, configs.Instance.EnvCloudAPIKey, configs.Instance.EnvCloudAPISecret)
	if err != nil {

		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: configs.Instance.EnvCloudUploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
