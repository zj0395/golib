package golog

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	dlog "github.com/rs/zerolog/log"
)

type LogConf struct {
	File  string
	Level int8
	Split bool // split warning/fatal and debug/info
}

type Logger struct {
	logger *zerolog.Logger
	fWarn  *os.File
	fPlain *os.File

	Conf *LogConf
}

// default logger
var defaultLogger *Logger = &Logger{
	logger: &dlog.Logger,
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(-1)
}

func GetDefault() *Logger {
	return defaultLogger
}

func SetDefault(l *Logger) {
	defaultLogger = l
}

// RedirectStdLog redirect log.* to logger
func RedirectStdLog() {
	log.SetOutput(defaultLogger.logger)
}

// NewLogger new by conf
func NewLogger(conf *LogConf) *Logger {
	var (
		fPlain, fWarn *os.File
		err           error
		logger        zerolog.Logger
	)

	fPlain, err = os.OpenFile(conf.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic("init log error" + err.Error())
	}

	// split warning/fatal to .wf
	if conf.Split {
		fWarn, err = os.OpenFile(conf.File+".wf", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic("init log error" + err.Error())
		}

		// min level: warn
		warnWriter := zerolog.MultiLevelWriter(fWarn)
		warnFilteredWriter := &minFilteredWriter{warnWriter, zerolog.WarnLevel}

		// max level: info
		plainWriter := zerolog.MultiLevelWriter(fPlain)
		plainFilteredWriter := &maxFilteredWriter{plainWriter, zerolog.InfoLevel}

		// combine to a logger
		w := zerolog.MultiLevelWriter(plainFilteredWriter, warnFilteredWriter)
		logger = zerolog.New(w)
	} else {
		logger = zerolog.New(fPlain)
	}

	logger = logger.With().Caller().Timestamp().Logger()
	logger.Level(zerolog.Level(conf.Level))

	return &Logger{
		logger: &logger,
		fPlain: fPlain,
		fWarn:  fWarn,
		Conf:   conf,
	}
}

// Close release fd
func (t *Logger) Close() {
	t.fPlain.Close()
	t.fWarn.Close()
	t.fPlain = nil
	t.fWarn = nil
}

func (t *Logger) Rename(suffix string) {
	fileName := t.Conf.File
	os.Rename(fileName, fileName+"."+suffix)
	if t.Conf.Split {
		os.Rename(fileName+".wf", fileName+".wf."+suffix)
	}
}

func (t *Logger) With() zerolog.Context {
	return t.logger.With()
}

// Trace level: -1
func (t *Logger) Trace() *zerolog.Event {
	return t.logger.Trace()
}

// Debug level: 0
func (t *Logger) Debug() *zerolog.Event {
	return t.logger.Debug()
}

// Info level: 1
func (t *Logger) Info() *zerolog.Event {
	return t.logger.Info()
}

// Warn level: 2
func (t *Logger) Warn() *zerolog.Event {
	return t.logger.Warn()
}

// Error level: 3
func (t *Logger) Error() *zerolog.Event {
	return t.logger.Error()
}

// Fatal level: 4
func (t *Logger) Fatal() *zerolog.Event {
	return t.logger.Fatal()
}

// Printf can not get right caller, to be fix by github
func (t *Logger) Printf(format string, args ...interface{}) {
	t.logger.Printf(format, args...)
}

func Trace() *zerolog.Event {
	return defaultLogger.Trace()
}

func Debug() *zerolog.Event {
	return defaultLogger.Debug()
}

func Info() *zerolog.Event {
	return defaultLogger.Info()
}

func Warn() *zerolog.Event {
	return defaultLogger.Warn()
}

func Error() *zerolog.Event {
	return defaultLogger.Error()
}

func Fatal() *zerolog.Event {
	return defaultLogger.Fatal()
}

func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}
