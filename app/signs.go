package app

import (
	"os"
	"syscall"
	"fmt"
	"os/signal"
	"sync"
)

type SignAction int8
const (
	GracefulExit SignAction = iota + 1
)

var gloableSingal SignAction
var signHandler func(*os.Signal)
var concernSigns []os.Signal
var signs chan os.Signal

func WatchSignals(wg *sync.WaitGroup){
	signs = make(chan os.Signal)
	signal.Notify(signs, concernSigns...)
	go func() {
		for {
			select{
			case  signal,ok := <-signs:
				if !ok {
					panic("signs occurs error")
				}
				fmt.Println("catched ",signal)
				signHandler(&signal)
			}
		}
	}()
}

func ConcernSigns(sigs ...os.Signal){
	concernSigns = sigs
}

func SetSignHandler(handler func(sign *os.Signal)){
	signHandler = handler
}

func DefaultSignHandler(sign *os.Signal){
	if *sign == syscall.SIGHUP || *sign == syscall.SIGHUP || *sign == syscall.SIGSTOP  || *sign == syscall.SIGQUIT  || *sign == syscall.SIGTERM  || *sign == syscall.SIGINT {
		SignDefer(GracefulExit)
	}
}

func SignDefer(act SignAction){
	gloableSingal = act
}
