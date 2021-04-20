package golog

import (
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
	dlog "github.com/rs/zerolog/log"
)

type LogConf struct {
	File  string
	Level int8
	Split bool
}

// default logger
var defaultLogger *zerolog.Logger = &dlog.Logger
var fPlain, fWarn *os.File = nil, nil

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func GetDefault() *zerolog.Logger {
	return defaultLogger
}

func SetDefault(l *zerolog.Logger) {
	defaultLogger = l
}

// RedirectStdLog redirect log.* to logger
func RedirectStdLog() {
	log.SetOutput(defaultLogger)
}

func Init(conf *LogConf) *zerolog.Logger {
	zerolog.SetGlobalLevel(-1)
	var err error = nil
	fPlain, err = os.OpenFile(conf.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fWarn, err = os.OpenFile(conf.File+".wf", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// min level: warn
	warnWriter := zerolog.MultiLevelWriter(fWarn)
	warnFilteredWriter := &minFilteredWriter{warnWriter, zerolog.WarnLevel}

	// max level: info
	plainWriter := zerolog.MultiLevelWriter(fPlain)
	plainFilteredWriter := &maxFilteredWriter{plainWriter, zerolog.InfoLevel}

	// combine to a logger
	w := zerolog.MultiLevelWriter(plainFilteredWriter, warnFilteredWriter)
	logger := zerolog.New(w).With().Caller().Timestamp().Logger()
	logger.Level(zerolog.Level(conf.Level))
	logger.Info().Str("f", "f")

	return &logger
}

// Close release fd
func Close() {
	fPlain.Close()
	fWarn.Close()
}

// Trace level: -1
func Trace() *zerolog.Event {
	return defaultLogger.Trace()
}

// Debug level: 0
func Debug() *zerolog.Event {
	return defaultLogger.Debug()
}

// Info level: 1
func Info() *zerolog.Event {
	return defaultLogger.Info()
}

// Warn level: 2
func Warn() *zerolog.Event {
	return defaultLogger.Warn()
}

// Error level: 3
func Error() *zerolog.Event {
	return defaultLogger.Error()
}

// Fatal level: 4
func Fatal() *zerolog.Event {
	return defaultLogger.Fatal()
}
