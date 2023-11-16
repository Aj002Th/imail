package logs

import (
	"log/slog"
	"os"
)

// Init 日志初始化
func Init() {
	opts := slog.HandlerOptions{
		AddSource: true,
	}
	handler := slog.NewJSONHandler(os.Stdout, &opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}