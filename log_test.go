package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	sync := InitWithPath("./", "prod", true)
	defer sync()
	Info("demo4:", String("app", "start ok"),
		Int("major version", 3))
	Error("demo4:", String("app", "crash"),
		Int("reason", -1))
	SugaredWarnf("print sugared warnf err: %s", "xxxx")
}

func TestStdLogger(t *testing.T) {
	StdLogger().Info("demo4:", String("app", "start ok"),
		Int("major version", 3))
	StdLogger().Error("demo4:", String("app", "crash"),
		Int("reason", -1))
	StdLogger().SugaredWarnf("print sugared warnf err: %s", "xxxx")
}
