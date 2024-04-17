package config

import (
	"os"
	"time"

	clog "github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

type config = string

const (
	AppEnv     config = "APP_ENV"
	MongoURI   config = "MONGO_URI"
	RedisURL   config = "REDIS_URL"
	Host       config = "HOST"
	Port       config = "PORT"
	Expiration config = "SESSION_EXPIRATION"
)

var defaults = map[config]string{
	AppEnv:     "dev",
	MongoURI:   "mongodb://localhost:27017",
	RedisURL:   "",
	Host:       "localhost",
	Port:       "8085",
	Expiration: "1m",
}

var C = map[config]string{}

var log = clog.WithPrefix("CONFIG")

func init() {
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Error(err)
		}
	}

	for k, v := range defaults {
		C[k] = loadConfig(k, v)
	}
}

func loadConfig(c config, def string) string {
	v := os.Getenv(string(c))
	if v == "" {
		return def
	}
	return v
}

// GetAppEnv reads the ExpirationTime from the environment variables, parses it and returns it as a time.Duration
func GetExpirationTime() time.Duration {
	pd := time.Minute * 1
	d := C[Expiration]

	defer func() {
		log.Infof("session duration is %s", d)
	}()

	if d == "" {
		d = "1m"
		return pd
	}

	pd, err := time.ParseDuration(d)
	if err != nil {
		log.Warn("config.GetExpirationTime", "err", err)
		return pd
	}

	return pd
}

func GetListenAddr() string {
	return C[Host] + ":" + C[Port]
}
