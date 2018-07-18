![](http://ww1.sinaimg.cn/large/7c998145ly1fte3roqfhij205k05k3yb.jpg)



# tigo
## This is a tiny framework for golang.


[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/letcgo/tigo?status.png)](http://godoc.org/github.com/letcgo/tigo)  [![travis](https://travis-ci.org/letcgo/tigo.svg?branch=master)](https://travis-ci.org/letcgo/tigo) [![Go Report Card](https://goreportcard.com/badge/github.com/letcgo/tigo)](https://goreportcard.com/report/github.com/letcgo/tigo) [![coveralls](https://coveralls.io/repos/letcgo/tigo/badge.svg?branch=master&service=github)](https://coveralls.io/github/letcgo/tigo?branch=master) 
![QQ群](https://img.shields.io/:QQ%E7%BE%A4-828486848-blue.svg)

# Feature
- can custom producer
- can custom workers
- can custom signal handlers
- logger with info/notice/warn/err
- cli params helper
- event monitor

# Quick demo

```go
package main

import (
	"tigo/app"
	"strconv"
	"fmt"
	"time"
	"os"
	"errors"
)

func main()  {
	dispatcher := new(app.Dispatcher)
	//optional
	app.RegistryLogger(app.NewLogger())
	//optional
	app.RegistryMonitor(app.NewMonitor())
	dispatcher.Setup(func(pipeline chan<- *app.Task)  {
		go func() {
			for i:=0;i<=999;i++ {
				pipeline<-&app.Task{
					ID: strconv.Itoa(i),
				}
			}
		}()
	}, nil)
	//optional
	dispatcher.Workers(3)
	dispatcher.RegistryWorker(func(task *app.Task) error {
		time.Sleep(3 * time.Second)
		fmt.Println("哈哈，task:",*task)
		panic(errors.New("worker error "))
		return nil
	})
	
    //optional	
	app.SetSignHandler(func(s *os.Signal) {
		fmt.Println("my sign handler", *s)
		app.DefaultSignHandler(s)
	})
	app.Start(dispatcher)

}


```




