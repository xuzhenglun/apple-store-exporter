package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	config = flag.String("config", "config.yaml", "config file path (default config.yaml)")
)

type Config struct {
	Endpoint string        `json:"endpoint" yaml:"endpoint"`
	Timeout  time.Duration `json:"timeout" yaml:"timeout"`

	Interval time.Duration `json:"interval" yaml:"interval"`
	Shops    []string      `json:"shops" yaml:"shops"`
	Models   []string      `json:"models" yaml:"models"`
}

func NewConfig() (*Config, error) {
	file, err := os.ReadFile(*config)
	if err != nil {
		return nil, fmt.Errorf("read config file failed, %w", err)
	}

	res := Config{}
	if err = yaml.Unmarshal(file, &res); err != nil {
		return nil, fmt.Errorf("load config failed, %w",err)
	}

	res.Default()

	return &res, nil
}

func (c *Config) Default() {
	if len(c.Endpoint) == 0 {
		c.Endpoint = "https://www.apple.com.cn"
	}

	if c.Timeout.Seconds() <= 0 {
		c.Timeout = time.Second * 2
	}

	if c.Interval.Seconds() <= 0 {
		c.Interval = time.Second * 15
	}
}
