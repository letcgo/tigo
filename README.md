![](http://ww1.sinaimg.cn/large/7c998145ly1fte3roqfhij205k05k3yb.jpg)



# tigo
## This is a tiny framework for golang.


[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/letcgo/tigo?status.png)](http://godoc.org/github.com/letcgo/tigo)  [![travis](https://travis-ci.org/letcgo/tigo.svg?branch=master)](https://travis-ci.org/letcgo/tigo) [![Go Report Card](https://goreportcard.com/badge/github.com/letcgo/tigo)](https://goreportcard.com/report/github.com/letcgo/tigo) [![coveralls](https://coveralls.io/repos/letcgo/tigo/badge.svg?branch=master&service=github)](https://coveralls.io/github/letcgo/tigo?branch=master) 
![QQ群](https://img.shields.io/:QQ%E7%BE%A4-828486848-blue.svg)

# Feature
- customize producer
- customize workers
- customize signal handlers
- customize logger with info/notice/warn/err
- built in cli params helper
- customize producer/worker checker, e.g. health checker
- event monitor

# Quick demo

```go
//this is main.go
package main

import (
	"github.com/tigo/app"
	"strconv"
	"time"
	"os"
	"log"
)

func main()  {
	dispatcher := new(app.Dispatcher)
	f,err := os.OpenFile("app.log", os.O_RDWR | os.O_CREATE , 0755)
	app.CheckErr(err)
	app.RegistryLogger(app.NewLogger(f, "", log.LstdFlags))
	//app.RegistryLogger(app.NewLogger(os.Stdout, "", log.LstdFlags))
	dispatcher.SetupHandler(func(pipeline chan<- *app.Task)  {
		go func() {
			for i:=0;i<=999;i++ {
				pipeline<-&app.Task{
					ID: strconv.Itoa(i),
				}
			}
		}()
	})
	dispatcher.RegistryChecker(func(e interface{}){
		app.GetLogger().Info("my checker catch: %v\n ", e)
	})
	dispatcher.RegistryWorkerChecker(func(e interface{}){
		app.NewMonitor().Notify("notify Dispatcher error", e)
		app.GetLogger().Info("my worker checker catch:  ", e)
	})
	dispatcher.Workers(3)
	dispatcher.RegistryWorker(func(task *app.Task) error {
		time.Sleep(1 * time.Second)
		app.GetLogger().Info("哈哈，task:",*task)
		panic("worker mocked error")
		return nil
	})
	app.SetSignHandler(func(s *os.Signal) {
		app.GetLogger().Info("my sign handler", *s)
		app.DefaultSignHandler(s)
	})
	app.Start(dispatcher)

}


```


```bash
go get ./...
go run src/app/main.go  -c ./.env
```

