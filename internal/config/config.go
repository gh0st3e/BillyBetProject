package config

import (
	"fmt"
	"github.com/gh0st3e/BillyBetProject/internal/util"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"net"
)

type Config struct {
	Database Database
	Server   Server
}
type Database struct {
	Host    string `required:"true" split_word:"true"`
	Port    string `required:"true" split_words:"true"`
	Name    string `required:"true" split_words:"true"`
	Address string `required:"false"`
}
type Server struct {
	Host    string `required:"true" split_word:"true"`
	Port    string `required:"true" split_words:"true"`
	Address string `required:"false"`
}

func Init() (*Config, error) {
	var cfg = Config{}

	err := initDB(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "config.Init.initDB couldn't init database")
	}

	err = initServer(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "config.Init.initServer couldn't init server")
	}

	return &cfg, nil
}
func initDB(cfg *Config) error {
	err := envconfig.Process(util.Database, &cfg.Database)
	if err != nil {
		return errors.Wrap(err, "config.initDB.Process couldn't load database from .env file")
	}

	cfg.Database.Address = fmt.Sprintf("(%s)/%s", net.JoinHostPort(cfg.Database.Host, cfg.Database.Port), cfg.Database.Name)

	return nil
}
func initServer(cfg *Config) error {
	err := envconfig.Process(util.Server, &cfg.Server)
	if err != nil {
		return errors.Wrap(err, "config.initDB.Process couldn't load database from .env file")
	}

	cfg.Server.Address = net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)

	return nil
}
