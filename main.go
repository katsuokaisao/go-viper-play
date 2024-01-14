package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	config, err := loadConfig("config/yaml", "smtp", "yaml")
	if err != nil {
		panic(fmt.Errorf("error in loadConfig from yaml: %w", err))
	}

	fmt.Printf("%+v\n", config)

	config, err = loadConfig("config/toml", "smtp", "toml")
	if err != nil {
		panic(fmt.Errorf("error in loadConfig from toml: %w", err))
	}

	fmt.Printf("%+v\n", config)

	config, err = loadConfig("config/json", "smtp", "json")
	if err != nil {
		panic(fmt.Errorf("error in loadConfig from json: %w", err))
	}

	fmt.Printf("%+v\n", config)

	os.Setenv("SMTP_SERVER_NAME", "xxx.gmail.com")
	config, err = loadConfig("config/json", "smtp", "json")
	if err != nil {
		panic(fmt.Errorf("error in loadConfig from json: %w", err))
	}

	fmt.Printf("%+v\n", config)
}

type Config struct {
	Smtp Smtp
}

type Smtp struct {
	Server   string `mapstructure:"server_name"`
	Port     int
	Address  string
	Password string
}

func loadConfig(dir, fileName, fileType string) (*Config, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath(dir)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error in viper ReadInConfig: %w", err)
	}

	var cofig Config
	if err := v.Unmarshal(&cofig); err != nil {
		return nil, fmt.Errorf("error in viper Unmarshal: %w", err)
	}

	return &cofig, nil
}
