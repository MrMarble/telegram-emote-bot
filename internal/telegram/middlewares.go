package telegram

import (
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

// MinQueryLength is the minimum length of the query to be considered.
func MinQueryLength(length int) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			if len(c.Query().Text) < length {
				return nil
			}

			return next(c)
		}
	}
}

// CacheQuery is a middleware that caches the result of the handler.
func CacheQuery(store *cache.Cache) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			query := c.Query().Text
			if answer, found := store.Get(query); found {
				log.Debug().Str("Query", query).Msg("Cache hit")
				return c.Answer(answer.(*telebot.QueryResponse))
			}

			return next(c)
		}
	}
}
