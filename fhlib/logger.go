package fhlib

import (
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"github.com/zj0395/golib/golog"
	"github.com/zj0395/golib/utils"
)

func GetLogger(ctx *fasthttp.RequestCtx) *zerolog.Logger {
	raw := ctx.UserValue(LoggerFlag)
	if v, ok := raw.(*zerolog.Logger); ok {
		return v
	}
	logger := golog.GetDefault().With().Str("logid", GetLogId(ctx)).Logger()
	SetLogger(ctx, &logger)
	return &logger
}

func SetLogger(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetUserValue(LoggerFlag, data)
}

func GetLogId(ctx *fasthttp.RequestCtx) string {
	raw := ctx.UserValue(LogIdFlag)
	if raw != nil {
		if v, ok := raw.(string); ok {
			return v
		}
	}

	const requestIDHeader = "X-Request-Id"
	var logid string
	if v := ctx.Request.Header.Peek(requestIDHeader); v != nil {
		logid = string(v)
	} else {
		logid = utils.GenLogId()
	}
	SetLogId(ctx, logid)

	return logid
}

func SetLogId(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetUserValue(LogIdFlag, data)
}
