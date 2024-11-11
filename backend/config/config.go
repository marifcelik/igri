package config

import (
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	clog "github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

// C is the global config variable
var C config

type Env string

const (
	DevEnv  Env = "dev"
	ProdEnv Env = "prod"
)

type config struct {
	AppEnv       Env    `envDefault:"dev"`
	MongoURI     string `env:"MONGO_URI,expand" envDefault:"mongodb://localhost:27017"`
	MongoName    string `envDefault:"go-chat"`
	RedisURL     string `env:"REDIS_URL,expand"`
	Host         string `envDefault:"localhost"`
	Port         string `envDefault:"8085"`
	Expiration   string `envDefault:"15m"`
	SessionIDKey string `envDefault:"user"`

	DBKey struct {
		Users             string `envDefault:"users"`
		Messages          string `envDefault:"messages"`
		Conversations     string `envDefault:"conversations"`
		UserConversations string `envDefault:"user_conversations"`
	} `envPrefix:"DB_KEY_"`

	HeaderKey struct {
		Session string `envDefault:"X-Session"`
		Expiry  string `envDefault:"X-Session-Expiry"`
	} `envPrefix:"HEADER_KEY_"`
}

var log = clog.WithPrefix("CONFIG")

func init() {
	if os.Getenv("APP_ENV") != string(ProdEnv) {
		if err := godotenv.Load(); err != nil {
			log.Warn("using default config", "err", err)
		}
	}

	C = config{}
	opts := env.Options{UseFieldNameByDefault: true}
	if err := env.ParseWithOptions(&C, opts); err != nil {
		log.Fatal(err)
	}
}

// GetAppEnv reads the ExpirationTime from the environment variables, parses it and returns it as a time.Duration
func GetExpirationTime() time.Duration {
	pd := time.Minute * 1
	d := C.Expiration

	defer func() {
		log.Infof("session duration is %s", d)
	}()

	if d == "" {
		d = "1m"
		return pd
	}

	pd2, err := time.ParseDuration(d)
	if err != nil {
		log.Warn("config.GetExpirationTime", "err", err)
		return pd
	}

	return pd2
}

// TODO implement the rest of the function
func GetIdleTimeout() time.Duration {
	log.Warn("config.GetIdleTimeout is not implemented")
	return 0
}
