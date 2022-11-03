package gosfile

import (
	"github.com/firdasafridi/gocrypt"
)

// Config will help set the config
type Config struct {

	// TypeEncryption can choosed from aes, des, rc4, or custom
	TypeEncryption string

	// Key will put when
	Key string

	// CustomEncryption will put in when the type encryption is custom
	CustomEncryption gocrypt.GocryptOption

	// Supported mime
	SupportedMime []string
}

func (cfg *Config) validation() error {

	if cfg == nil {
		return ErrConfig
	}

	if len(cfg.SupportedMime) == 0 {
		cfg.SupportedMime = listSupportedMime
	}

	if cfg.TypeEncryption == CUSTOM {
		if cfg.CustomEncryption == nil {
			return ErrCustomEncMandatory
		}
		return nil
	}

	if len(cfg.Key) == 0 {
		return ErrKeyEmpty
	}
	return nil
}

func (cfg *Config) getEncryption() (gocrypt.GocryptOption, error) {
	switch cfg.TypeEncryption {

	case AES:
		aesOpt, err := gocrypt.NewAESOpt(cfg.Key)
		if err != nil {
			return nil, err
		}
		return aesOpt, nil

	case DES:
		desOpt, err := gocrypt.NewDESOpt(cfg.Key)
		if err != nil {
			return nil, err
		}
		return desOpt, nil

	case RC4:
		rc4, err := gocrypt.NewRC4Opt(cfg.Key)
		if err != nil {
			return nil, err
		}
		return rc4, nil
	}

	return cfg.CustomEncryption, nil
}
