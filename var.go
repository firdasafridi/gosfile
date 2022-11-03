package gosfile

import "errors"

const (
	AES    = "AES"
	DES    = "DES"
	RC4    = "RC4"
	CUSTOM = "CUSTOM"
)

var (
	ErrCustomEncMandatory = errors.New("custom encryption need to set CustomEncryption as mandatory option")
	ErrKeyEmpty           = errors.New("key is required")
	ErrConfig             = errors.New("config is required")

	ErrInternal     = errors.New("internal error")
	ErrDecodeBase64 = errors.New("err decode to base64")

	ErrMimeNotSupported = errors.New("not supported mime type")
)

var listSupportedMime = []string{
	"image/jpeg",
	"image/png",
	"image/bmp",
	"application/pdf",
	"image/svg+xml",
	"text/plain",
	"image/webp",
}
