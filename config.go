package govdev

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	Debug      bool
	Hidebanner bool
	Database   DatabaseConfig
	Cache      CacheConfig
}

type DatabaseConfig struct {
	Engine   string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	Host     string
}

type CacheConfig struct {
	Engine   string
	Host     string
	Port     string
	Password string
	DBName   int
}

func LoadConfig() (Config, []error) {
	var errs []error

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to find config file, error: %s\n", err.Error())
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Unable to unmarshal into struct, error: %s\n", err.Error())
	}
	return conf, errs
}
