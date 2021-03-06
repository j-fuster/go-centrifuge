package configstore

import "github.com/centrifuge/go-centrifuge/errors"

const (
	// ErrConfigStorageBootstrap must be returned when there is an error while bootstrapping the config storage
	ErrConfigStorageBootstrap = errors.Error("error when bootstrapping config storage")
)
