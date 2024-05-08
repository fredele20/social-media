package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultServerPort   = "3000"
	defaultConfigName   = "dev"
	defaultDatabasePort = 27017
	configType          = "env"
	currentDir          = "."
	// defaultApiVersion             = "1.0.0"
	// defaultEnableWalletConversion = false
)

var envVars = []string{
	"ENV",
	"DATABASE_URL",
	"DATABASE_NAME",
	"JWT_SECRETKEY",
}

type Config struct {
	Env          string `mapstructure:"ENV"`
	DatabaseURL  string `mapstructure:"DATABASE_URL"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	JwtSecretKey string `mapstructure:"JWT_SECRETKEY"`
	HttpPort     string `mapstructure:"HTTP_PORT"`
}

func Load() (*Config, error) {
	configName := defaultEnv("ENV", defaultConfigName)

	viper.AddConfigPath(currentDir)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AutomaticEnv()

	for _, v := range envVars {
		if err := viper.BindEnv(v); err != nil {
			return nil, err
		}
	}

	viper.SetDefault("HTTP_PORT", defaultPort(configName))
	viper.SetDefault("DATABASE_PORT", defaultDatabasePort)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, skipping...")
		} else {
			return nil, err
		}
	}

	c := Config{}
	err = viper.Unmarshal(&c)

	return &c, err
}

// var s Secrets

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("error loading .env file")
// 	}

// 	s = Secrets{}

// 	s.DatabaseURL = os.Getenv("DATABASE_URL")
// 	s.DatabaseName = os.Getenv("DATABASE_NAME")
// 	s.JwtSecretKey = os.Getenv("JWT_SECRETKEY")

// 	if s.HttpPort = os.Getenv("HTTP_PORT"); s.HttpPort == "" {
// 		s.HttpPort = "90"
// 	}
// }

// func GetSecrets() Secrets {
// 	return s
// }

func defaultPort(env string) string {
	defaultPort := defaultServerPort
	if env != "dev" {
		os.Unsetenv("HTTP_PORT")
		defaultPort = defaultEnv("PORT", defaultServerPort)
	}
	return defaultPort
}

func defaultEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}
