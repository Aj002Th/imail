package logs

import (
	"github.com/Aj002Th/imail/common/crontab"
	"github.com/natefinch/lumberjack"
	"log/slog"
)

var (
	Logger *slog.Logger

	logFileManager *lumberjack.Logger

	logFileName = "imail"
	logFileExt  = ".log"
	cron        = "0 0 * * *"
)

// Init 日志初始化
func Init() {
	// 日志文件管理
	logFileManager = &lumberjack.Logger{
		Filename:   "./data/logs/" + logFileName + logFileExt,
		MaxSize:    500, // megabytes
		MaxBackups: 10,
		MaxAge:     30, // days
		LocalTime:  true,
	}

	// 定时进行日志切割
	err := crontab.StartScheduledTasks(cron, func() {
		if err := logFileManager.Rotate(); err != nil {
			slog.Error(err.Error())
		}
	})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	// slog 日志初始化
	opts := slog.HandlerOptions{
		AddSource: true,
	}
	handler := slog.NewJSONHandler(logFileManager, &opts)
	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}

func Close() {
	logFileManager.Close()
}
