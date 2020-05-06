package common

import (
	"errors"

	"github.com/mattermost/mattermost-server/v5/model"
)

type KVStore interface {
	Load(key string) ([]byte, error)
	Store(key string, data []byte) error
	StoreTTL(key string, data []byte, ttlSeconds int64) error
	StoreWithOptions(key string, value []byte, opts model.PluginKVSetOptions) (bool, error)
	Delete(key string) error
}

var ErrNotFound = errors.New("not found")
