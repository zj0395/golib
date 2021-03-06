package ginlog

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zj0395/golib/golog"
)

// GetGinLogger Log for gin
//	r := gin.New()
//	r.Use(golog.GetGinLogger("gin"))
func GetGinLogger(serName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		// after request
		latency := time.Since(t)
		// clientIP := c.ClientIP()
		// method := c.Request.Method
		// statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		golog.Info().Str("ser_name", serName).Str("method", c.Request.Method).Str("path", path).Dur("resp_time", latency).Int("status", c.Writer.Status()).Str("client_ip", c.ClientIP()).Msg(msg)
	}
}
