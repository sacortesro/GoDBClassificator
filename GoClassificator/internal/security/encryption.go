package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"

	"GoClassificator/internal/logger"
)

// Encrypt encrypts a text using AES
func Encrypt(text string) (string, error) {
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(key) != 32 {
		return "", errors.New("invalid encryption key length, must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts an AES-encrypted text
func Decrypt(encryptedText string) (string, error) {
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(key) != 32 {
		logger.AppLogger.ErrorLog.Println("Invalid encryption key length, must be 32 bytes")
		return "", errors.New("invalid encryption key length, must be 32 bytes")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
