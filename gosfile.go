package gosfile

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/firdasafridi/gocrypt"
)

// GosFile contains functions for file encryption.
type GosFile struct {
	option            gocrypt.GocryptOption
	listSupportedMime []string
}

// New creates a new GosFile instance with the given configuration.
// It returns an error if the configuration validation fails or if the encryption option cannot be retrieved.
func New(cfg *Config) (*GosFile, error) {

	if err := cfg.validation(); err != nil {
		return nil, err
	}

	cryptOption, err := cfg.getEncryption()
	if err != nil {
		return nil, err
	}
	return &GosFile{
		option:            cryptOption,
		listSupportedMime: cfg.SupportedMime,
	}, nil
}

// EncryptFromPath encrypts a file from the given path.
// It returns an Option containing the encrypted data or an error.
func (gosFile *GosFile) EncryptFromPath(path string) *Option {

	fileB, err := os.ReadFile(path)
	if err != nil {
		return newOption("", err)
	}

	return gosFile.EncryptFromByte(fileB)
}

// EncryptFromByte encrypts a byte slice of file data.
// It returns an Option containing the encrypted data or an error.
func (gosFile *GosFile) EncryptFromByte(fileB []byte) *Option {

	mimeType := http.DetectContentType(fileB)

	if !gosFile.isSupportedMime(mimeType) {
		return newOption("", ErrMimeNotSupported)
	}

	encodeString := fmt.Sprintf("data:%s;base64,%s", mimeType,
		string(base64.StdEncoding.EncodeToString(fileB)))

	return gosFile.EncryptFromBase64(encodeString)
}

// EncryptFromBase64 encrypts a base64 encoded string.
// It returns an Option containing the encrypted data or an error.
func (gosFile *GosFile) EncryptFromBase64(encodeString string) *Option {

	encString, err := gosFile.option.Encrypt([]byte(encodeString))
	if err != nil {
		return newOption("", ErrMimeNotSupported)
	}

	return newOption(encString, nil)
}

// isSupportedMime checks if the given MIME type is supported for encryption.
// It returns true if supported, otherwise false.
func (gosFile *GosFile) isSupportedMime(mimeContent string) bool {
	for _, mimeType := range gosFile.listSupportedMime {
		if mimeContent == mimeType {
			return true
		}
	}
	return false
}

// EncryptFromFile is a placeholder function, it currently returns nil and should be implemented to encrypt a file from the given path.
func (gosFile *GosFile) EncryptFromFile(path string) *Option {

	return nil
}

// DecryptFromPath decrypts encrypted data from a file at the given path.
// It returns an Option containing the decrypted data or an error.
func (gosFile *GosFile) DecryptFromPath(path string) *Option {

	encByte, err := os.ReadFile(path)
	if err != nil {
		return newOption("", err)
	}

	return gosFile.DecryptFromByte(encByte)
}

// DecryptFromByte decrypts a byte slice of encrypted data.
// It returns an Option containing the decrypted data or an error.
func (gosFile *GosFile) DecryptFromByte(encByte []byte) *Option {

	decString, err := gosFile.option.Decrypt(encByte)
	if err != nil {
		return newOption("", ErrMimeNotSupported)
	}

	return newOption(decString, nil)
}
