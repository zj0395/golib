package fhlib

import (
	"github.com/valyala/fasthttp"

	"github.com/zj0395/golib/liberr"
)

const (
	ErrorFlag  = "errorflag"
	DataFlag   = "dataflag"
	CustomFlag = "customflag"

	LoggerFlag = "loggerflag"
	LogIdFlag  = "logidflag"
)

func GetError(ctx *fasthttp.RequestCtx) error {
	raw := ctx.UserValue(ErrorFlag)
	if raw == nil {
		return nil
	}
	if v, ok := raw.(error); ok {
		return v
	}
	return liberr.ErrorFlagValError
}

func SetError(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetUserValue(ErrorFlag, err)
}

func GetData(ctx *fasthttp.RequestCtx) interface{} {
	return ctx.UserValue(DataFlag)
}

func SetData(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetUserValue(DataFlag, data)
}

func GetCustom(ctx *fasthttp.RequestCtx) interface{} {
	return ctx.UserValue(CustomFlag)
}

func SetCustom(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetUserValue(CustomFlag, data)
}
