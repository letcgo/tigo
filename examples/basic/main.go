package main

import (
	"github.com/letcgo/tigo"
	"strconv"
	"time"
	"os"
	"log"
)

func main()  {
	dispatcher := new(tigo.Dispatcher)
	f,err := os.OpenFile("app.log", os.O_RDWR | os.O_CREATE , 0755)
	tigo.CheckErr(err)
	tigo.RegistryLogger(tigo.NewLogger(f, "", log.LstdFlags))
	//app.RegistryLogger(app.NewLogger(os.Stdout, "", log.LstdFlags))
	dispatcher.SetupHandler(func(pipeline chan<- *tigo.Task)  {
		go func() {
			for i:=0;i<=999;i++ {
				pipeline<-&tigo.Task{
					ID: strconv.Itoa(i),
				}
			}
		}()
	})
	dispatcher.RegistryChecker(func(e interface{}){
		tigo.GetLogger().Info("my checker catch: %v\n ", e)
	})
	dispatcher.RegistryWorkerChecker(func(e interface{}){
		tigo.NewMonitor().Notify("notify Dispatcher error", e)
		tigo.GetLogger().Info("my worker checker catch:  ", e)
	})
	dispatcher.Workers(3)
	dispatcher.RegistryWorker(func(task *tigo.Task) error {
		time.Sleep(1 * time.Second)
		tigo.GetLogger().Info("哈哈，task:",*task)
		panic("worker mocked error")
		return nil
	})
	tigo.SetSignHandler(func(s *os.Signal) {
		tigo.GetLogger().Info("my sign handler", *s)
		tigo.DefaultSignHandler(s)
	})
	tigo.Start(dispatcher)
}

