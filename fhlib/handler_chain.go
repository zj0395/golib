package fhlib

import (
	"runtime/debug"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/zj0395/golib/liberr"
)

// control exec of handlers
type wrapper struct {
	middleWares []fasthttp.RequestHandler
	final       fasthttp.RequestHandler
}

func NewWrapper(h fasthttp.RequestHandler) *wrapper {
	return &wrapper{
		middleWares: []fasthttp.RequestHandler{},
		final:       h,
	}
}

func (t *wrapper) Add(handler fasthttp.RequestHandler) *wrapper {
	t.middleWares = append(t.middleWares, handler)
	return t
}

func (t *wrapper) Exec(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()

	logger := GetLogger(ctx)

	defer func() {
		if p := recover(); p != nil {
			fatalErr := string(debug.Stack())
			logger.Fatal().Str("stack", fatalErr).Msg("[FATAL 500]")
			SetErrorOutput(ctx, liberr.PanicError)
		}
		logger.Info().Int64("costms", time.Since(startTime).Milliseconds()).Msg("Request Done")
	}()

	for _, h := range t.middleWares {
		h(ctx)
		// stop if already error
		if err := GetError(ctx); err != nil {
			SetErrorOutput(ctx, err)
			return
		}
	}
	if t.final != nil {
		t.final(ctx)
	}
	if err := GetError(ctx); err != nil {
		SetErrorOutput(ctx, err)
	} else if data := GetData(ctx); data != nil {
		SetOutput(ctx, data)
	}
}
