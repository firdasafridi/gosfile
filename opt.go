package gosfile

import (
	"encoding/base64"
	"os"
	"strings"
)

type Option struct {
	src string
	err error
}

func newOption(src string, err error) *Option {
	return &Option{
		src: src,
		err: err,
	}
}

func (opt *Option) ToString() (string, error) {
	if err := opt.validate(); err != nil {
		return "", err
	}

	return string(opt.src), nil
}

func (opt *Option) ExportTo(pathAndName string) error {
	if err := opt.validate(); err != nil {
		return err
	}

	listWord := strings.Split(string(opt.src), "base64,")
	if len(listWord) != 2 {
		return opt.writeFile([]byte(opt.src), pathAndName)
	}

	wordBase64 := listWord[1]
	decodeString, err := base64.StdEncoding.DecodeString(wordBase64)
	if err != nil {
		return err
	}

	return opt.writeFile(decodeString, pathAndName)
}

func (opt *Option) writeFile(wordBase64 []byte, pathAndName string) error {

	fPlain, err := os.Create(pathAndName)
	if err != nil {
		return err
	}
	defer fPlain.Close()

	if _, err := fPlain.Write(wordBase64); err != nil {
		return err
	}

	if err := fPlain.Sync(); err != nil {
		return err
	}
	return nil
}

func (opt *Option) validate() (err error) {
	if opt == nil {
		return ErrInternal
	}

	if opt.err != nil {
		return err
	}

	return nil
}
