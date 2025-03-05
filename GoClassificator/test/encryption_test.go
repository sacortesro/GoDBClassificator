package test

import (
	"fmt"
	"os"
	"testing"

	"GoClassificator/internal/config"
	"GoClassificator/internal/security"

	"github.com/stretchr/testify/assert"
)

// Set the environment variables based on config file
func SetLoadEnv() {
	config.LoadEnv()
}

// TestEncryptDecrypt tests the Encrypt and Decrypt functions
func TestEncryptDecrypt(t *testing.T) {
	// os.Setenv("ENCRYPTION_KEY", "thisis32bitlongpassphraseimusing")
	SetLoadEnv()
	originalText := "Hello, World!"

	encryptedText, err := security.Encrypt(originalText)
	fmt.Println(encryptedText)
	assert.NoError(t, err)
	assert.NotEmpty(t, encryptedText)

	decryptedText, err := security.Decrypt(encryptedText)
	fmt.Println(decryptedText)
	assert.NoError(t, err)
	assert.Equal(t, originalText, decryptedText)
}

// TestEncryptInvalidKeyLength tests the Encrypt function with invalid key length
func TestEncryptInvalidKeyLength(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "shortkey")
	_, err := security.Encrypt("Hello, World!")
	assert.Error(t, err)
	assert.Equal(t, "invalid encryption key length, must be 32 bytes", err.Error())
}

// TestDecryptInvalidKeyLength tests the Decrypt function with invalid key length
func TestDecryptInvalidKeyLength(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "shortkey")
	_, err := security.Decrypt("someencryptedtext")
	assert.Error(t, err)
	assert.Equal(t, "invalid encryption key length, must be 32 bytes", err.Error())
}
