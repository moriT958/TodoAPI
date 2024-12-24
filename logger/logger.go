package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	Log  *slog.Logger
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})
}
