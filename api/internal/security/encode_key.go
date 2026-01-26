package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

func deriveKey(secret string) []byte {
	hash := sha256.Sum256([]byte(secret))
	return hash[:]
}

func EncryptKey(plainText string, secretKey string) (string, error) {
	key := deriveKey(secretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", e.GenerateError(*EncryptError, err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", e.GenerateError(*EncryptError, err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", e.GenerateError(*EncryptError, err)
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptKey(cipherTextBase64 string, secretKey string) (string, error) {
	key := deriveKey(secretKey)

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", e.GenerateError(*DecryptError, err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", e.GenerateError(*DecryptError, err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", e.GenerateError(*DecryptError, err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", e.GenerateError(*DecryptError, errors.New("texto cifrado muito curto"))
	}

	nonce, cipherTextClean := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherTextClean, nil)
	if err != nil {
		return "", e.GenerateError(*DecryptError, err)
	}

	return string(plainText), nil
}
