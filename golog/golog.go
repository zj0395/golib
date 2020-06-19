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
	Level zerolog.Level
}

// default logger
var logger zerolog.Logger = dlog.Logger
var fPlain, fWarn *os.File = nil, nil

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Init(conf *LogConf) {
	zerolog.SetGlobalLevel(conf.Level)

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

	log.SetOutput(fPlain)

	// min level: warn
	warnWriter := zerolog.MultiLevelWriter(fWarn)
	warnFilteredWriter := &minFilteredWriter{warnWriter, zerolog.WarnLevel}

	// max level: info
	plainWriter := zerolog.MultiLevelWriter(fPlain)
	plainFilteredWriter := &maxFilteredWriter{plainWriter, zerolog.InfoLevel}

	// combine to a logger
	w := zerolog.MultiLevelWriter(plainFilteredWriter, warnFilteredWriter)
	logger = zerolog.New(w).With().Caller().Timestamp().Logger()
}

// Close release fd
func Close() {
	fPlain.Close()
	fWarn.Close()
}

// Trace level: -1
func Trace() *zerolog.Event {
	return logger.Trace()
}

// Debug level: 0
func Debug() *zerolog.Event {
	return logger.Debug()
}

// Info level: 1
func Info() *zerolog.Event {
	return logger.Info()
}

// Warn level: 2
func Warn() *zerolog.Event {
	return logger.Warn()
}

// Error level: 3
func Error() *zerolog.Event {
	return logger.Error()
}

// Fatal level: 4
func Fatal() *zerolog.Event {
	return logger.Fatal()
}
