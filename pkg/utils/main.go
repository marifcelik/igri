package utils

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func init() {
	if GetAppEnv() != "prod" {
		os.Setenv("APP_ENV", "dev")
		if err := godotenv.Load(); err != nil {
			log.Error(err)
		}
	}
}

func GetAppEnv() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}
	return appEnv
}

func GetDBURL() string {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "ws://localhost:8080"
	}
	return connStr
}

// if REDIS_URL env is not set then that returns "" because 'redis.Config' defaults
// are same the default connection url
func GetRedisURL() string {
	connStr := os.Getenv("REDIS_URL")
	return connStr
}

func GetListenAddr() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}
	// we don't check host because if it's empty, that's fine
	return host + ":" + port
}

func GetExpirationTime() time.Duration {
	d := os.Getenv("SESSION_EXPIRATION")
	defaultD := time.Hour * 24
	log := log.WithPrefix("SESSION")

	if d == "" {
		log.Info("session duration is 24h")
		return defaultD
	}

	pd, err := time.ParseDuration(d)
	if err != nil {
		log.Warn("utils.GetExpirationTime", "err", err)
		log.Info("session duration is 24h")
		return defaultD
	}

	log.Infof("session duration is %s", d)
	return pd
}

// Check the error and exit if its not nil.
// The parameters after the second parameter will be joined into a single string
func CheckErr(err error, msgParams ...string) {
	msg := strings.Join(msgParams, "")

	if err != nil {
		if msg != "" {
			log.Fatal(msg, "err", err)
		} else {
			log.Fatal(err)
		}
	}
}
