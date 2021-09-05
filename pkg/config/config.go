package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init(configPath string) (*Config, error) {
	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgresql.Username = os.Getenv("POSTGRES_USER")
	cfg.Postgresql.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgresql.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgresql.Port = os.Getenv("POSTGRES_PORT")

	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgresql); err != nil {
		return err
	}

	return nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}
