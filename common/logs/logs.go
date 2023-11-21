package logs

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// Init 日志初始化
func Init() {
	opts := slog.HandlerOptions{
		AddSource: true,
	}
	handler := slog.NewJSONHandler(os.Stdout, &opts)
	Logger = slog.New(handler)

	slog.SetDefault(Logger)
}
