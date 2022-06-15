package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

type BotConfig struct {
	Token           string        `env:"TOKEN" env-required:""`
	CacheExpiration time.Duration `env:"CACHE_EXPIRATION" env-default:"5m"`
	CachePurge      time.Duration `env:"CACHE_PURGE" env-default:"10m"`
	TelegramCache   int           `env:"TELEGRAM_CACHE" env-default:"300"`
}

var cfg BotConfig

func LoadConfig() BotConfig {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	return cfg
}

func GetConfig() BotConfig {
	return cfg
}
