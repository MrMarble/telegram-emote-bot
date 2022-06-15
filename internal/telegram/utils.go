package telegram

import (
	"github.com/mrmarble/telegram-emote-bot/pkg/betterttv"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

func emoteToResult(emotes []betterttv.SearchResult) tele.Results {
	results := make(tele.Results, len(emotes))

	for i, emote := range emotes {
		switch mime := emote.Type; {
		case mime == "gif":
			results[i] = &tele.GifResult{
				URL:       emote.ID.Url(),
				Title:     emote.Code + " - " + emote.Type,
				ThumbURL:  emote.ID.Url(),
				ThumbMIME: "image/gif",
				Cache:     string(emote.ID),
			}

			break
		case mime == "png":
			results[i] = &tele.PhotoResult{
				URL:         emote.ID.Url(),
				Title:       emote.Code,
				ThumbURL:    emote.ID.Url(),
				Description: emote.Code + " - " + emote.Type,
				Cache:       string(emote.ID),
			}

			break
		default:
			log.Warn().Str("type", mime).Msg("Unknown mime type")
		}

		results[i].SetResultID(string(emote.ID))
	}

	return results
}
