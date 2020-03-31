package golog

import (
	"github.com/rs/zerolog"
)

type minFilteredWriter struct {
	w     zerolog.LevelWriter
	level zerolog.Level
}

func (w *minFilteredWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)

}
func (w *minFilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= w.level {
		return w.w.WriteLevel(level, p)
	}

	return len(p), nil
}

type maxFilteredWriter struct {
	w     zerolog.LevelWriter
	level zerolog.Level
}

func (w *maxFilteredWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)

}
func (w *maxFilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level <= w.level {
		return w.w.WriteLevel(level, p)
	}

	return len(p), nil
}
