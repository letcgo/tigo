package app

import (
	"sync"
	"fmt"
	"log"
)

type TaskGenerator func(chan<- *Task)
type DispatchChecker func()error

type IDispatch interface {
	Dispatch(*sync.WaitGroup)
}


type Dispatcher struct {
	ID            int
	workerNum     int
	taskGenerator TaskGenerator
	checker       DispatchChecker
	backlogNum int

	taskHandler TaskHandler
	workChecker WorkChecker

	logger *log.Logger
}

var parent *Dispatcher


func Start(app IDispatch)  {
	wg := &sync.WaitGroup{}
	WatchSignals(wg)
	app.Dispatch(wg)
	wg.Wait()
	println("app exit normally")
}


func (i *Dispatcher)Backlog(num int){
	i.backlogNum = num
}
func (i *Dispatcher)Setup(handler TaskGenerator, checker DispatchChecker){
	i.taskGenerator = handler
	i.checker = checker
}


func (i *Dispatcher) RegistryWorker(handler TaskHandler){
	i.taskHandler = handler
}
func (i *Dispatcher) RegistryChecker(checker WorkChecker){
	i.workChecker = checker
}


func (i *Dispatcher) Workers(consumerNum int){
	i.workerNum = consumerNum
}

func (i *Dispatcher)Dispatch(wg *sync.WaitGroup){
	defer func(){
		if err := recover(); nil != err {
			monitor.Notify("Dispatcher error", err)
			logger.Err("Dispatcher error", err)
			i.Dispatch(wg)
		}
	}()
	var task *Task
	pipeline := make(chan *Task, i.backlogNum)
	i.taskGenerator(pipeline)

	if nil == signHandler {
		signHandler = DefaultSignHandler
	}
	var workers []*Worker
	for it:=0; it< i.workerNum; it++  {
		wg.Add(1)
		worker := new(Worker)
		workers = append(workers, worker)
		go func(worker *Worker) {
			defer wg.Done()
			defer func() {
				if err := recover(); nil != err {
					fmt.Println("catched ", err)
				}
				println("worker done ")
			}()
			for{
				select {
				case task,_ = <-pipeline:
					worker.handler = i.taskHandler
					worker.workChecker = i.workChecker
					worker.work(task)

					if gloableSingal == GracefulExit {
						return
					}
				}
			}
		}(worker)
	}
}

