package cloudsigma

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetLogin() (Config, error) {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, err
}
