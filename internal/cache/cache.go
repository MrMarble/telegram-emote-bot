package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
)

var store *cache.Cache

func innitCache(expiration, cleanupInterval time.Duration) {
	once := sync.Once{}
	once.Do(func() {
		log.Debug().Msg("New cache initialized")
		store = cache.New(expiration, cleanupInterval)
	})
}

func New(expiration, cleanupInterval time.Duration) *cache.Cache {
	if store == nil {
		innitCache(expiration, cleanupInterval)
	}

	return store
}
