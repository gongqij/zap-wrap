package log

import (
	"math/rand"
	"testing"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestDebug(t *testing.T) {
	t.Parallel()
	loggerSync := Init("tmp", true)
	defer loggerSync()
	for i := 0; i < 20000; i++ {
		DebugWithFields("TestDebug", ZapString("RandStr", RandStringRunes(10)))
	}
}

func TestSugaredDebugf(t *testing.T) {
	t.Parallel()
	sync := Init("tmp", true)
	defer sync()
	for i := 0; i < 20000; i++ {
		Debugf("TestSugaredDebug:RandStr=%s", RandStringRunes(10))
	}
}

func TestInfo(t *testing.T) {
	t.Parallel()
	sync := Init("tmp", true)
	defer sync()
	for i := 0; i < 20000; i++ {
		InfoWithFields("TestInfo", ZapString("RandStr", RandStringRunes(10)))
	}
}

func TestSugaredInfof(t *testing.T) {
	t.Parallel()
	sync := Init("tmp", true)
	defer sync()
	for i := 0; i < 20000; i++ {
		Infof("TestSugaredInfo:RandStr=%s", RandStringRunes(10))
	}
}
