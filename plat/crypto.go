package plat

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func HmacSha256Signer(message string, secretKey string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(mac.Sum(nil))
}
