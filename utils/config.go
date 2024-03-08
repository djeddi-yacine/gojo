package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Environment          string        // ENVIRONMENT
	DBSource             string        // DB_SOURCE
	HTTPServerAddress    string        // HTTP_SERVER_ADDRESS
	GRPCServerAddress    string        // GRPC_SERVER_ADDRESS
	RedisQueueAddress    string        // REDIS_QUEUE_ADDRESS
	RedisCacheAddress    string        // REDIS_CACHE_ADDRESS
	MeilisearchAddress   string        // MEILISEATCH_ADDRESS
	TokenSymmetricKey    string        // TOKEN_SYMMETRIC_KEY
	MeiliSearchMasterKey string        // MEILISEATCH_MASTER_KEY
	AccessTokenDuration  time.Duration // ACCESS_TOKEN_DURATION
	RefreshTokenDuration time.Duration // REFRESH_TOKEN_DURATION
	CacheRepetition      uint8         // CACHE_REPETITION
	CacheKeyDuration     time.Duration // CACHE_KEY_DURATION
	CacheCountDuration   time.Duration // CACHE_COUNT_DURATION
	EmailSenderName      string        // EMAIL_SENDER_NAME
	EmailSenderAddress   string        // EMAIL_SENDER_ADDRESS
	EmailSenderPassword  string        // EMAIL_SENDER_PASSWORD
}

func LoadConfig(path string, name string) (*Config, error) {
	file, err := os.Open(path + "/" + name + ".env")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	var (
		key   string
		value string
	)
	re := regexp.MustCompile(`'[^']*'`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key = strings.TrimSpace(parts[0])
		matches := re.FindAllString(parts[1], -1)
		for _, match := range matches {
			value = strings.TrimSpace(match[1 : len(match)-1])
		}

		switch key {
		case "ENVIRONMENT":
			config.Environment = value
		case "DB_SOURCE":
			config.DBSource = value
		case "HTTP_SERVER_ADDRESS":
			config.HTTPServerAddress = value
		case "GRPC_SERVER_ADDRESS":
			config.GRPCServerAddress = value
		case "REDIS_QUEUE_ADDRESS":
			config.RedisQueueAddress = value
		case "REDIS_CACHE_ADDRESS":
			config.RedisCacheAddress = value
		case "MEILISEATCH_ADDRESS":
			config.MeilisearchAddress = value
		case "TOKEN_SYMMETRIC_KEY":
			config.TokenSymmetricKey = value
		case "MEILISEATCH_MASTER_KEY":
			config.MeiliSearchMasterKey = value
		case "ACCESS_TOKEN_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing ACCESS_TOKEN_DURATION: %v", err)
			}
			config.AccessTokenDuration = duration
		case "REFRESH_TOKEN_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing REFRESH_TOKEN_DURATION: %v", err)
			}
			config.RefreshTokenDuration = duration
		case "CACHE_REPETITION":
			repetition, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return nil, err
			}
			config.CacheRepetition = uint8(repetition)
		case "CACHE_KEY_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing CACHE_KEY_DURATION: %v", err)
			}
			config.CacheKeyDuration = duration
		case "CACHE_COUNT_DURATION":
			duration, err := time.ParseDuration(strings.TrimSpace(value))
			if err != nil {
				return nil, fmt.Errorf("error parsing CACHE_COUNT_DURATION: %v", err)
			}
			config.CacheCountDuration = duration
		case "EMAIL_SENDER_NAME":
			config.EmailSenderName = value
		case "EMAIL_SENDER_ADDRESS":
			config.EmailSenderAddress = value
		case "EMAIL_SENDER_PASSWORD":
			config.EmailSenderPassword = value
		}
	}

	// Set environment variables
	os.Setenv("ENVIRONMENT", config.Environment)
	os.Setenv("DB_SOURCE", config.DBSource)
	os.Setenv("HTTP_SERVER_ADDRESS", config.HTTPServerAddress)
	os.Setenv("GRPC_SERVER_ADDRESS", config.GRPCServerAddress)
	os.Setenv("REDIS_QUEUE_ADDRESS", config.RedisQueueAddress)
	os.Setenv("REDIS_CACHE_ADDRESS", config.RedisCacheAddress)
	os.Setenv("MEILISEATCH_ADDRESS", config.MeilisearchAddress)
	os.Setenv("TOKEN_SYMMETRIC_KEY", config.TokenSymmetricKey)
	os.Setenv("MEILISEATCH_MASTER_KEY", config.MeiliSearchMasterKey)
	os.Setenv("ACCESS_TOKEN_DURATION", config.AccessTokenDuration.String())
	os.Setenv("REFRESH_TOKEN_DURATION", config.RefreshTokenDuration.String())
	os.Setenv("CACHE_REPETITION", strconv.Itoa(int(config.CacheRepetition)))
	os.Setenv("CACHE_KEY_DURATION", config.CacheKeyDuration.String())
	os.Setenv("CACHE_COUNT_DURATION", config.CacheCountDuration.String())
	os.Setenv("EMAIL_SENDER_NAME", config.EmailSenderName)
	os.Setenv("EMAIL_SENDER_ADDRESS", config.EmailSenderAddress)
	os.Setenv("EMAIL_SENDER_PASSWORD", config.EmailSenderPassword)

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return config, nil
}
