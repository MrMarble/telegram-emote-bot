package telegram

import (
	"time"

	ch "github.com/mrmarble/telegram-emote-bot/internal/cache"
	"github.com/mrmarble/telegram-emote-bot/internal/config"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

type Telegram struct {
	bot   *telebot.Bot
	cache *cache.Cache
	cgf   config.BotConfig

	handlersRegistered    bool
	middlewaresRegistered bool
}

func New(cfg config.BotConfig) (*Telegram, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, _ telebot.Context) {
			log.Error().Str("module", "telegram").Err(err).Msg("telebot internal error")
		},
	})
	if err != nil {
		return nil, err
	}

	log.Info().Str("module", "telegram").Int64("id", bot.Me.ID).Str("name", bot.Me.FirstName).Str("username", bot.Me.Username).Msg("connected to telegram")

	return &Telegram{
		bot:   bot,
		cache: ch.GetCache(cfg.CacheExpiration, cfg.CachePurge),
		cgf:   cfg,
	}, nil
}

// Start starts polling for telegram updates
func (t *Telegram) Start() {
	t.RegisterMiddlewares()
	t.registerHandlers()

	log.Info().Str("module", "telegram").Msg("start polling")
	t.bot.Start()
}

// Stop stops polling for telegram updates
func (t *Telegram) Stop() {
	log.Info().Str("module", "telegram").Msg("stop polling")
	t.bot.Stop()
}

// RegisterMiddlewares registers all middlewares
func (t *Telegram) RegisterMiddlewares() {
	if t.middlewaresRegistered {
		return
	}

	log.Info().Str("module", "telegram").Msg("registering middlewares")

	t.bot.Use(MinQueryLength(3))
	t.bot.Use(CacheQuery(t.cache))

	t.middlewaresRegistered = true
}

// RegisterHandlers registers all the handlers
func (t *Telegram) registerHandlers() {
	if t.handlersRegistered {
		return
	}

	log.Info().Str("module", "telegram").Msg("registering handlers")

	t.bot.Handle(telebot.OnQuery, t.handleQuery)

	t.handlersRegistered = true
}
