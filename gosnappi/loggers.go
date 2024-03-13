package gosnappi

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	Logger      *zerolog.Logger
	logLevel    zerolog.Level = zerolog.InfoLevel
	logToFile   bool          = false
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

func SetLogger(usrDefinedLogger zerolog.Logger) {
	Logger = &usrDefinedLogger
}

func SetLogOutputToFile(choice bool) zerolog.Logger {
	logToFile = choice
	return getLogger(ModuleName)
}

func SetLogLevel(level zerolog.Level) {
	logLevel = level
	refreshLogLevel(logLevel)
}

func SetLogFileName(fileName string) {
	logFileName = fileName
}

func getLogger(name string) zerolog.Logger {
	if ModuleName == "" {
		if err := initlog(); err != nil {
			panic(fmt.Errorf("Logger helper set failed: %v", err))
		}
	}
	ModuleName = name
	if logFileName == "" {
		logFileName = name
	}
	var localLogger zerolog.Logger
	if !logToFile {
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().In(time.Local)
		}
		zerolog.TimeFieldFormat = time.RFC3339Nano
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-1-02T15:04:05.000Z"}
		localLogger = zerolog.New(consoleWriter).Level(logLevel).With().Timestamp().Str("Module", name).Logger()

	} else {
		if err := initFileLogger(); err != nil {
			panic(fmt.Errorf("Logger init failed: %v", err))
		}
		localLogger = log.With().Str("Module", name).Logger()
	}
	Logger = &localLogger
	return *Logger
}

func initLogger() (err error) {
	writer := lumberjack.Logger{
		Filename:   path.Join(*logSt.logDir, fmt.Sprintf("%s.log", logFileName)),
		MaxSize:    *logSt.maxLogSizeMB,
		MaxBackups: *logSt.maxLogBackups,
	}
	zerolog.TimestampFieldName = "ts"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = zerolog.New(&writer).Level(logLevel).With().Timestamp().Logger()
	return err
}

func refreshLogLevel(logLevel zerolog.Level) {
	if !logToFile {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(logLevel)
	} else {
		log.Logger = log.Level(logLevel)
	}
}

func initFileLogger() (err error) {
	if err = os.RemoveAll(*logSt.logDir); err != nil {

		return err
	}
	if err = os.MkdirAll(*logSt.logDir, os.ModePerm); err != nil {
		return err
	}
	if err := os.Chmod(*logSt.logDir, 0771); err != nil {
		return err
	}
	return initLogger()
}
