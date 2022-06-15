package telegram

import (
	"github.com/mrmarble/telegram-emote-bot/pkg/betterttv"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

func (t *Telegram) handleQuery(ctx telebot.Context) error {
	query := ctx.Query().Text
	log.Debug().Str("Query", query).Msg("Query received")

	emotes, err := betterttv.DefaultClient.Search(query)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to search for emotes")
	}

	log.Debug().Str("Query", query).Int("Emotes", len(emotes)).Msg("Search finished")

	results := emoteToResult(emotes)
	response := &telebot.QueryResponse{
		Results:   results,
		CacheTime: t.cgf.TelegramCache,
	}

	t.cache.Set(query, response, 0)

	return ctx.Answer(response)
}
