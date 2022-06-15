package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mrmarble/telegram-emote-bot/internal/config"
	"github.com/mrmarble/telegram-emote-bot/internal/telegram"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	log.Info().Msg("Starting bot")

	cfg := config.LoadConfig()

	bot, err := telegram.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bot")
	}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	go func(gracefulStop chan os.Signal) {
		<-gracefulStop

		defer bot.Stop()

		log.Info().Msg("Stopping bot")
	}(gracefulStop)

	bot.Start()
}
