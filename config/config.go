package config

import (
	"errors"
	"github.com/companieshouse/gofigure"
	log "github.com/sirupsen/logrus"
	"sync"
)

// Config holds configuration details set by the environment
type Config struct {
	TransactionsMongoDBURL      string `env:"TRANSACTIONS_MONGODB_URL"       flag:"transactions-mongodb-url"       flagDesc:"Transactions MongoDB server URL"`
	TransactionsMongoDBDatabase string `env:"TRANSACTIONS_MONGODB_DATABASE"  flag:"transactions-mongodb-database"  flagDesc:"Transactions MongoDB database for data"`
	LogLevel                    string `env:"LOG_LEVEL"                      flag:"log-level"                      flagDesc:"Logging level of the application"`
}

var cfg *Config
var mtx sync.Mutex

// Get returns a pointer to a Config instance
// populated with values from environment or command-line flags
func Get() (*Config, error) {

	mtx.Lock()
	defer mtx.Unlock()

	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{}

	err := gofigure.Gofigure(cfg)
	if err != nil {
		return nil, err
	}

	mandatoryConfigsMissing := false

	if cfg.TransactionsMongoDBURL == "" {
		log.Warn("TRANSACTIONS_MONGODB_URL not set in environment")
		mandatoryConfigsMissing = true
	}

	if cfg.TransactionsMongoDBDatabase == "" {
		log.Warn("TRANSACTIONS_MONGODB_DATABASE not set in environment")
		mandatoryConfigsMissing = true
	}

	if mandatoryConfigsMissing {
		return nil, errors.New("mandatory configs missing from environment")
	}

	return cfg, nil
}
