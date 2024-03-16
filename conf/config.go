package conf

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Environment string // ENVIRONMENT
	DB          *DBEnv
	Server      *ServerEnv
	Data        *DataEnv
	Email       *EmailEnv
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		config = &Config{
			DB:     &DBEnv{},
			Server: &ServerEnv{},
			Data:   &DataEnv{},
			Email:  &EmailEnv{},
		}
		scanner = bufio.NewScanner(file)
		x       = regexp.MustCompile(`\s*#.*$|^\s+|\s+$`)
	)

	for scanner.Scan() {
		txt := scanner.Text()
		if !strings.Contains(txt, "=") {
			continue
		}

		line := strings.TrimSpace(x.ReplaceAllString(txt, ""))
		kv := strings.SplitN(line, "=", 2)
		if len(kv) < 1 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := kv[1]

		switch key {
		case "ENVIRONMENT":
			config.Environment = value
		case "DB_NAME":
			config.DB.NAME = value
		case "DB_HOST":
			config.DB.HOST = value
		case "DB_USER":
			config.DB.USER = value
		case "DB_PASSWORD":
			config.DB.PASSWORD = value
		case "DB_PORT":
			port, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("error parsing DB_PORT: %w", err)
			}
			config.DB.PORT = port
		case "HTTP_ADDRESS":
			addr, err := url.Parse("http://" + value)
			if err != nil {
				return nil, fmt.Errorf("error parsing HTTP_ADDRESS: %w", err)
			}
			config.Server.HTTPAddress = addr
		case "GRPC_ADDRESS":
			addr, err := url.Parse("http://" + value)
			if err != nil {
				return nil, fmt.Errorf("error parsing GRPC_ADDRESS: %w", err)
			}
			config.Server.GRPCAddress = addr
		case "REDIS_QUEUE_ADDRESS":
			addr, err := url.Parse("redis://" + value)
			if err != nil {
				return nil, fmt.Errorf("error parsing REDIS_QUEUE_ADDRESS: %w", err)
			}
			config.Server.RedisQueueAddress = addr
		case "REDIS_CACHE_ADDRESS":
			addr, err := url.Parse("redis://" + value)
			if err != nil {
				return nil, fmt.Errorf("error parsing REDIS_CACHE_ADDRESS: %w", err)
			}
			config.Server.RedisCacheAddress = addr
		case "MEILISEATCH_ADDRESS":
			addr, err := url.Parse("http://" + value)
			if err != nil {
				return nil, fmt.Errorf("error parsing MEILISEATCH_ADDRESS: %w", err)
			}
			config.Server.MeilisearchAddress = addr
		case "TOKEN_SYMMETRIC_KEY":
			config.Data.TokenSymmetricKey = value
		case "MEILISEATCH_MASTER_KEY":
			config.Data.MeiliSearchMasterKey = value
		case "ACCESS_TOKEN_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing ACCESS_TOKEN_DURATION: %w", err)
			}
			config.Data.AccessTokenDuration = duration
		case "REFRESH_TOKEN_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing REFRESH_TOKEN_DURATION: %w", err)
			}
			config.Data.RefreshTokenDuration = duration
		case "CACHE_REPETITION":
			repetition, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return nil, err
			}
			config.Data.CacheRepetition = uint8(repetition)
		case "CACHE_KEY_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing CACHE_KEY_DURATION: %w", err)
			}
			config.Data.CacheKeyDuration = duration
		case "CACHE_COUNT_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing CACHE_COUNT_DURATION: %w", err)
			}
			config.Data.CacheCountDuration = duration
		case "EMAIL_SENDER_NAME":
			config.Email.Name = value
		case "EMAIL_SENDER_ADDRESS":
			config.Email.Address = value
		case "EMAIL_SENDER_PASSWORD":
			config.Email.Password = value
		}
	}

	// Set environment variables
	os.Setenv("ENVIRONMENT", config.Environment)
	os.Setenv("DB_NAME", config.DB.NAME)
	os.Setenv("DB_HOST", config.DB.HOST)
	os.Setenv("DB_USER", config.DB.USER)
	os.Setenv("DB_PASSWORD", config.DB.PASSWORD)
	os.Setenv("DB_PORT", fmt.Sprint(config.DB.PORT))
	os.Setenv("HTTP_ADDRESS", config.Server.HTTPAddress.Host)
	os.Setenv("GRPC_ADDRESS", config.Server.GRPCAddress.Host)
	os.Setenv("REDIS_QUEUE_ADDRESS", config.Server.RedisQueueAddress.Host)
	os.Setenv("REDIS_CACHE_ADDRESS", config.Server.RedisCacheAddress.Host)
	os.Setenv("MEILISEATCH_ADDRESS", config.Server.MeilisearchAddress.Host)
	os.Setenv("TOKEN_SYMMETRIC_KEY", config.Data.TokenSymmetricKey)
	os.Setenv("MEILISEATCH_MASTER_KEY", config.Data.MeiliSearchMasterKey)
	os.Setenv("ACCESS_TOKEN_DURATION", config.Data.AccessTokenDuration.String())
	os.Setenv("REFRESH_TOKEN_DURATION", config.Data.RefreshTokenDuration.String())
	os.Setenv("CACHE_REPETITION", strconv.Itoa(int(config.Data.CacheRepetition)))
	os.Setenv("CACHE_KEY_DURATION", config.Data.CacheKeyDuration.String())
	os.Setenv("CACHE_COUNT_DURATION", config.Data.CacheCountDuration.String())
	os.Setenv("EMAIL_SENDER_NAME", config.Email.Name)
	os.Setenv("EMAIL_SENDER_ADDRESS", config.Email.Address)
	os.Setenv("EMAIL_SENDER_PASSWORD", config.Email.Password)

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return config, nil
}
