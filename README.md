![](http://ww1.sinaimg.cn/large/7c998145ly1fte3roqfhij205k05k3yb.jpg)



# tigo
## This is a tiny framework for golang.

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


