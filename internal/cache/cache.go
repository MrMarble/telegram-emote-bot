package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
)

var (
	store *cache.Cache
	once  sync.Once
)

// GetCache returns the cache store
func GetCache(expiration, cleanupInterval time.Duration) *cache.Cache {
	once.Do(func() {
		log.Debug().Msg("New cache initialized")
		store = cache.New(expiration, cleanupInterval)
	})

	return store
}
