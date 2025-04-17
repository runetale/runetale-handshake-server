package utility

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"golang.org/x/exp/slog"
)

const (
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

const (
	JsonFmtStr string = "json"
	TextFmtStr string = "text"
)

const (
	StdOutFilePath string = "/dev/stdout"
	StdErrFilePath string = "/dev/stderr"
)

var (
	logger  *zap.Logger
	devMode bool = false
)

type Logger struct {
	*slog.Logger
}

func NewLogger(file *os.File, logFmt, logLevel string) (*Logger, error) {
	logger, err := initSlog(file, logFmt, logLevel)
	if err != nil {
		return nil, err
	}

	return &Logger{
		Logger: logger,
	}, nil
}

func initSlog(file *os.File, logFmt, logLevel string) (*slog.Logger, error) {
	var level slog.Level
	switch logLevel {
	case DebugLevelStr:
		level = slog.LevelDebug
	case InfoLevelStr:
		level = slog.LevelInfo
	case WarningLevelStr:
		level = slog.LevelWarn
	case ErrorLevelStr:
		level = slog.LevelError
	default:
		return nil, fmt.Errorf("unknown log level %s", logLevel)
	}

	opts := slog.HandlerOptions{
		Level: level,
	}

	switch logFmt {
	case JsonFmtStr:
		return slog.New(opts.NewJSONHandler(file)), nil
	case TextFmtStr:
		return slog.New(opts.NewTextHandler(file)), nil
	default:
		return nil, fmt.Errorf("unknown log format %s", logFmt)
	}
}
