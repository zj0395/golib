package golog

import (
	"sync"
	"testing"
	"time"
)

func TestRotateLogLost(t *testing.T) {
	file := "log/tmp.log"
	conf := &RotateLogConf{
		LogConf: LogConf{
			File:  file,
			Level: -1,
			Split: false,
		},

		EnableRotate: true,
		Period:       1,
		UseMinute:    true,
	}
	logger := NewRotateLog(conf)
	SetDefault(logger.GetLogger())

	var wg sync.WaitGroup
	routeCnt := 500
	runCnt := 20000
	for i := 0; i < routeCnt; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < runCnt; j++ {
				Info().Int("t1", j).Int("routine", i).Msg("helloworld")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	logger.Close()

	// check log lost
}

func BenchmarkPrint(b *testing.B) {
	conf := &RotateLogConf{
		LogConf: LogConf{
			File:  "log/bench.log",
			Level: -1,
		},

		EnableRotate: true,
		Period:       1,
		UseMinute:    true,
	}
	logger := NewRotateLog(conf)
	defer logger.Close()
	SetDefault(logger.GetLogger())

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info().
				Str("string", "four!").
				Time("time", time.Time{}).
				Int("int", 123).
				Float32("float", -2.203230293249593).
				Msg("hello world")
		}
	})
}
