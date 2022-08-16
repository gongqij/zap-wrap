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
		Debug("TestDebug", ZapString("RandStr", RandStringRunes(10)))
	}
}

func TestSugaredDebug(t *testing.T) {
	t.Parallel()
	sync := Init("tmp", true)
	defer sync()
	for i := 0; i < 20000; i++ {
		SugaredDebugf("TestSugaredDebug:RandStr=%s", RandStringRunes(10))
	}
}

func TestInfo(t *testing.T) {
	t.Parallel()
	sync := Init("tmp", true)
	defer sync()
	for i := 0; i < 20000; i++ {
		Info("TestInfo", ZapString("RandStr", RandStringRunes(10)))
	}
}

func TestSugaredInfo(t *testing.T) {
	t.Parallel()
	sync := Init("dev", true)
	defer sync()
	for i := 0; i < 20000; i++ {
		SugaredInfof("TestSugaredInfo:RandStr=%s", RandStringRunes(10))
	}
}
