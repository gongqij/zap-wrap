package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ContextLogKey = "__context_log_key__"
)

const (
	requestStartTimeKey   = "start"
	requestEndTimeKey     = "end"
	requestDurationKey    = "duration"
	requestStatusKey      = "status"
	requestURLRawQueryKey = "query"
	requestClientIPKey    = "ip"
	requestUserAgentKey   = "user-agent"
	requestURLPath        = "path"

	nameKey = "name"
)

// GinHandler return a gin middleware, you should add it to gin server.
func GinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		fields := []Field{
			ZapString(requestStartTimeKey, start.Format(logTimeFormatter)),
			ZapString(requestURLPath, path),
		}
		var fieldsInterface []interface{}
		fieldsInterface = append(fieldsInterface, requestStartTimeKey, start.Format(logTimeFormatter))
		fieldsInterface = append(fieldsInterface, requestURLPath, path)
		c.Set(ContextLogKey, stdLogger.withFields(fields, fieldsInterface))
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				stdLogger.Error(e)
			}
		} else {
			fields := []Field{
				ZapInt(requestStatusKey, c.Writer.Status()),
				ZapString(requestStartTimeKey, start.Format(logTimeFormatter)),
				ZapString(requestEndTimeKey, end.Format(logTimeFormatter)),
				ZapDuration(requestDurationKey, latency),
				ZapString(requestURLRawQueryKey, c.Request.URL.RawQuery),
				ZapString(requestClientIPKey, c.ClientIP()),
				ZapString(requestUserAgentKey, c.Request.UserAgent()),
			}
			stdLogger.Info(c.Request.Method+" "+path, fields...)
		}
	}
}

// GetFromGin return a Entry, you should use Entry in api handler to log
func GetFromGin(c *gin.Context) *Logger {
	var (
		e  *Logger
		ok bool
	)
	ee, ok := c.Get(ContextLogKey)
	if ok {
		e, ok = ee.(*Logger)
	}
	if ok {
		return e
	} else {
		return stdLogger.withFields(nil, nil)
	}
}

// GetFromGinWithName is same as GetFromGin and you can pass a name to it, then log will add a field called name
func GetFromGinWithName(c *gin.Context, name string) *Logger {
	e := GetFromGin(c)
	logger := e.logger.With(ZapString(nameKey, name))
	sugaredLogger := e.sugaredLogger.With(nameKey, name)
	e = &Logger{logger, sugaredLogger}
	c.Set(ContextLogKey, e)

	return e
}
