package gosnappi

import (
	"fmt"
	"os"
	"path"

	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

type logInfo struct {
	rootDir       *string
	logDir        *string
	maxLogSizeMB  *int
	maxLogBackups *int
}

var (
	logSt       *logInfo
	Logger      *slog.Logger
	logLevel    slog.Leveler = slog.LevelWarn
	logToFile   bool         = false
	logFileName string
	ModuleName  string
)

func initlog() error {
	logSt = &logInfo{
		rootDir:       new(string),
		logDir:        new(string),
		maxLogSizeMB:  new(int),
		maxLogBackups: new(int),
	}
	*logSt.maxLogSizeMB = 25
	*logSt.maxLogBackups = 5
	if *logSt.rootDir = os.Getenv("SRC_ROOT"); *logSt.rootDir == "" {
		*logSt.rootDir = "."
	}
	*logSt.logDir = path.Join(*logSt.rootDir, "logs")
	Logger = nil
	ModuleName = ""
	return nil
}

func SetLogOutputToFile(choice bool) {
	logToFile = choice
}

func SetLogLevel(level slog.Leveler) {
	logLevel = level
}

func SetLogFileName(fileName string) {
	logFileName = fileName
}

func getLogger(name string) slog.Logger {
	if ModuleName == "" {
		if err := initlog(); err != nil {
			panic(fmt.Errorf("Logger helper set failed: %v", err))
		}
	}
	ModuleName = name
	if logFileName == "" {
		logFileName = name
	}
	var localLogger *slog.Logger
	var err error
	if !logToFile {
		handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: logLevel,
		})
		localLogger = slog.New(handler).With("Module", name)
	} else {
		localLogger, err = initFileLogger(name)
		if err != nil {
			panic(fmt.Errorf("Logger init failed: %v", err))

		}
	}
	Logger = localLogger
	return *Logger
}

func initLogger(name string) *slog.Logger {
	writer := &lumberjack.Logger{
		Filename:   path.Join(*logSt.logDir, fmt.Sprintf("%s.log", logFileName)),
		MaxSize:    *logSt.maxLogSizeMB,
		MaxBackups: *logSt.maxLogBackups,
	}
	// slog requires a handler. We'll use JSONHandler writing to lumberjack
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: logLevel,
	})

	logger := slog.New(handler).With("Module", name)
	return logger
}

func initFileLogger(name string) (*slog.Logger, error) {
	if err := os.RemoveAll(*logSt.logDir); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(*logSt.logDir, os.ModePerm); err != nil {
		return nil, err
	}
	if err := os.Chmod(*logSt.logDir, 0771); err != nil {
		return nil, err
	}
	return initLogger(name), nil
}
