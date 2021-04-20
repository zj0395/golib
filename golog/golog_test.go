package golog

import (
	"sync"
	"testing"
)

func TestPrint(t *testing.T) {
	conf := LogConf{
		File:  "tmp.log",
		Level: -1,
	}
	SetDefault(Init(&conf))

	var wg sync.WaitGroup
	routeCnt := 1000
	for i := 0; i < routeCnt; i++ {
		wg.Add(1)
		go func(i int) {
			runCnt := 1000
			for j := 0; j < runCnt; j++ {
				Info().Int("t1", j).Int("routine", i).Msg("helloworld")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	Close()
}
