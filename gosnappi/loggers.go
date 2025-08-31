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

type logger struct {
	logSt       *logInfo
	logger      *slog.Logger
	logLevel    slog.Level
	levelVar    *slog.LevelVar
	logToFile   bool
	logFileName string
	moduleName  string
}

type LoggerInterface interface {
	SetLogOutputToFile(bool) LoggerInterface
	SetLogFileName(string) LoggerInterface
	SetLogLevel(slog.Level) LoggerInterface
}

func (l *logger) initlog() error {
	l.logSt = &logInfo{
		rootDir:       new(string),
		logDir:        new(string),
		maxLogSizeMB:  new(int),
		maxLogBackups: new(int),
	}
	*l.logSt.maxLogSizeMB = 25
	*l.logSt.maxLogBackups = 5
	if *l.logSt.rootDir = os.Getenv("SRC_ROOT"); *l.logSt.rootDir == "" {
		*l.logSt.rootDir = "."
	}
	*l.logSt.logDir = path.Join(*l.logSt.rootDir, "logs")
	if l.levelVar == nil {
		l.levelVar = new(slog.LevelVar)
	}
	l.logger = nil
	l.logLevel = slog.LevelWarn
	l.moduleName = ""
	return nil
}

func (l *logger) getLogger(name string) slog.Logger {
	if l.moduleName == "" {
		if err := l.initlog(); err != nil {
			panic(fmt.Errorf("Logger helper set failed: %v", err))
		}
	}
	l.moduleName = name
	if l.logFileName == "" {
		l.logFileName = name
	}
	var localLogger *slog.Logger
	var err error
	if !l.logToFile {
		l.levelVar.Set(l.logLevel)
		handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: l.levelVar,
		})
		localLogger = slog.New(handler).With("Module", name)
	} else {
		localLogger, err = l.initFileLogger(name)
		if err != nil {
			panic(fmt.Errorf("Logger init failed: %v", err))

		}
	}
	l.logger = localLogger
	return *l.logger
}

func (l *logger) initLogger(name string) *slog.Logger {
	writer := &lumberjack.Logger{
		Filename:   path.Join(*l.logSt.logDir, fmt.Sprintf("%s.log", l.logFileName)),
		MaxSize:    *l.logSt.maxLogSizeMB,
		MaxBackups: *l.logSt.maxLogBackups,
	}
	// slog requires a handler. We'll use JSONHandler writing to lumberjack
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: l.levelVar,
	})

	logger := slog.New(handler).With("Module", name)
	return logger
}

func (l *logger) initFileLogger(name string) (*slog.Logger, error) {
	if err := os.RemoveAll(*l.logSt.logDir); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(*l.logSt.logDir, os.ModePerm); err != nil {
		return nil, err
	}
	if err := os.Chmod(*l.logSt.logDir, 0771); err != nil {
		return nil, err
	}
	return l.initLogger(name), nil
}

// This instruncts the logging module to log to a File and not to console.
// Please Note if you want to change the log file name first call SetLogFileName and then this fucntion to start logging to a file
func (l *logger) SetLogOutputToFile(choice bool) LoggerInterface {
	l.logToFile = choice
	logs = l.getLogger(l.moduleName)
	return l
}

// Set the log level for logging by default the log level is set to warnings
func (l *logger) SetLogLevel(level slog.Level) LoggerInterface {
	l.logLevel = level
	if l.levelVar == nil {
		l.levelVar = new(slog.LevelVar)
	}
	l.levelVar.Set(level)
	return l
}

// The file name with which log is generated , please note the file will be present under the logs directory
// By default the file name is same as the name of the module
func (l *logger) SetLogFileName(fileName string) LoggerInterface {
	l.logFileName = fileName
	return l
}

// Provides you with the gosnappi logger functionality to customise logging
func Logger() LoggerInterface {
	return loggerSt
}
