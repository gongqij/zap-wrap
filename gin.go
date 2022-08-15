package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ContextLogKey = "__context_log_key__"
)

const (
	requestStartTimeKey = "start"
	queryPathKey        = "path"
	nameKey             = "name"
)

// GinHandler return a gin middleware, you should add it to gin server.
func GinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := []Field{
			ZapString(requestStartTimeKey, time.Now().Format(logTimeFormatter)),
			ZapString(queryPathKey, c.Request.URL.Path),
		}
		var fieldsInterface []interface{}
		fieldsInterface = append(fieldsInterface, requestStartTimeKey, time.Now().Format(logTimeFormatter))
		fieldsInterface = append(fieldsInterface, queryPathKey, c.Request.URL.Path)
		c.Set(ContextLogKey, stdLogger.withFields(fields, fieldsInterface))
		c.Next()
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
