package gosfile

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/firdasafridi/gocrypt"
)

// GosFile will contain function encryption file
type GosFile struct {
	option            gocrypt.GocryptOption
	listSupportedMime []string
}

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

func (gosFile *GosFile) EncryptFromPath(path string) *Option {

	fileB, err := os.ReadFile(path)
	if err != nil {
		return newOption("", err)
	}

	return gosFile.EncryptFromByte(fileB)
}

func (gosFile *GosFile) EncryptFromByte(fileB []byte) *Option {

	mimeType := http.DetectContentType(fileB)

	if !gosFile.isSupportedMime(mimeType) {
		return newOption("", ErrMimeNotSupported)
	}

	encodeString := fmt.Sprintf("data:%s;base64,%s", mimeType,
		string(base64.StdEncoding.EncodeToString(fileB)))

	return gosFile.EncryptFromBase64(encodeString)
}

func (gosFile *GosFile) EncryptFromBase64(encodeString string) *Option {

	encString, err := gosFile.option.Encrypt([]byte(encodeString))
	if err != nil {
		return newOption("", ErrMimeNotSupported)
	}

	return newOption(encString, nil)
}

func (gosFile *GosFile) isSupportedMime(mimeContent string) bool {
	for _, mimeType := range gosFile.listSupportedMime {
		if mimeContent == mimeType {
			return true
		}
	}
	return false
}

func (gosFile *GosFile) EncryptFromFile(path string) *Option {

	return nil
}

func (gosFile *GosFile) DecryptFromPath(path string) *Option {

	encByte, err := os.ReadFile(path)
	if err != nil {
		return newOption("", err)
	}

	return gosFile.DecryptFromByte(encByte)
}

func (gosFile *GosFile) DecryptFromByte(encByte []byte) *Option {

	decString, err := gosFile.option.Decrypt(encByte)
	if err != nil {
		return newOption("", ErrMimeNotSupported)
	}

	return newOption(decString, nil)
}
