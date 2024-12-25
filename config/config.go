package config

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	cfg  *config
	once sync.Once
)

type config struct {
	Version      string `json:"Version"`
	Address      string `json:"Address"`
	ReadTimeout  int64  `json:"ReadTimeout"`
	WriteTimeout int64  `json:"WriteTimeout"`
}

func Load(filename string) error {
	once.Do(func() {
		cfg = new(config)
	})

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return err
	}

	return nil
}

var (
	Version      = &cfg.Version
	Address      = &cfg.Address
	ReadTimeout  = &cfg.ReadTimeout
	WriteTimeout = &cfg.WriteTimeout
)
