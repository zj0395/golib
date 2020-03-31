package golog

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	conf := LogConf{
		"tmp.log",
		1,
	}
	Trace().Msg("hello world")
	Init(&conf)
	Trace().Msg("hello world")
	Info().Msg("hello world")
	time.Sleep(10 * time.Second)
	Debug().Msg("hello world")
	Warn().Msg("hello world")
	Close()
}
