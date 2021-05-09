package golog

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type RotateConf struct {
}

type RotateLogConf struct {
	LogConf

	EnableRotate bool
	Backup       int  // log file save days before delete
	Period       int  // unit: depend on UseMinute
	UseMinute    bool // true: period mean minute, false: period is hour
}

type RotateLog struct {
	*Logger
	Conf       *RotateLogConf
	rotateCh   chan bool
	cancelFunc func()
}

func NewRotateLog(conf *RotateLogConf) *RotateLog {
	if !conf.EnableRotate {
		return &RotateLog{
			Logger: NewLogger(&conf.LogConf),
		}
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	logger := &RotateLog{
		Logger: NewLogger(&conf.LogConf),

		Conf:       conf,
		rotateCh:   make(chan bool),
		cancelFunc: cancelFunc,
	}

	period := time.Hour * time.Duration(conf.Period)
	if conf.UseMinute {
		period = time.Minute * time.Duration(conf.Period)
	}

	go func() {
		for {
			now := time.Now()
			nextRotateTime := now.Truncate(period).Add(period).Add(time.Second)
			timer := time.NewTimer(nextRotateTime.Sub(now))

			select {
			case <-timer.C:
				logger.rotateCh <- true
			case <-ctx.Done():
				close(logger.rotateCh)
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-logger.rotateCh:
				t := time.Now().Add(-time.Second * 20)
				suffix := fmt.Sprintf("%04d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour())
				if conf.UseMinute {
					suffix += strconv.Itoa(t.Minute())
				}
				logger.Rename(suffix)

				var oldLogger *Logger
				oldLogger, logger.Logger = logger.Logger, NewLogger(&logger.Conf.LogConf)
				SetDefault(logger.Logger)

				// wait buffer flush, then close
				// in a http request, a handler may `getDefault`, then
				time.Sleep(10 * time.Second)
				oldLogger.Close()

				//			go deleteExpiredLog(period)
			}
		}
	}()

	return logger
}

func (t *RotateLog) GetLogger() *Logger {
	return t.Logger
}

func (t *RotateLog) Close() {
	t.cancelFunc()
	t.Logger.Close()
}
