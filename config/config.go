package config

import (
	"encoding/json"
	"just-do-it/logger"
	"os"
	"sync"
)

type config struct {
	Version string `json:"version"`
	Address string `json:"address"`
}

var (
	Cfg  *config
	once sync.Once
)

func InitConfig(cfgFile string) error {
	once.Do(func() {
		Cfg = new(config)

	})

	file, err := os.Open(cfgFile)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&Cfg); err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}
