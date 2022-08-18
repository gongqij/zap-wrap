package log

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestGinLog(t *testing.T) {

	loggerSync := Init("ginlog", true)
	defer loggerSync()

	testServer := gin.New()
	testServer.Use(GinHandler(), RecoveryWithZap(true))

	testServer.GET("/ginlog", func(c *gin.Context) {
		time.Sleep(time.Second * 2)
		log := GetFromGin(c)
		log.Info("ginlog test")
		panic("ginlog test panic")
	})
	testServer.GET("/ginlogwithname", func(c *gin.Context) {
		time.Sleep(time.Second)
		log := GetFromGinWithName(c, "test")
		log.Info("ginlog test with name")
	})

	go testServer.Run(":18080")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		res, err := http.Get("http://localhost:18080/ginlog")
		if err != nil {
			t.Errorf("http GET err:%s", err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(string(body))
		}
	}()

	go func() {
		defer wg.Done()
		res, err := http.Get("http://localhost:18080/ginlogwithname")
		if err != nil {
			t.Errorf("http GET err:%s", err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(string(body))
		}
	}()
	wg.Wait()
}
