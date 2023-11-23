package logs

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

var Logger *slog.Logger

// Init 日志初始化
func Init() {
	opts := slog.HandlerOptions{
		AddSource: true,
	}

	// todo: close file
	logFilename := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile("./data/logs/"+logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	//handler := slog.NewJSONHandler(os.Stdout, &opts)
	handler := slog.NewJSONHandler(file, &opts)
	Logger = slog.New(handler)

	slog.SetDefault(Logger)
}
