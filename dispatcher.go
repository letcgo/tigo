package tigo

import (
	"sync"
	"log"
)

type TaskGenerator func(chan<- *Task)
type Checker func(interface{})

type IDispatch interface {
	Dispatch(*sync.WaitGroup)
}


type Dispatcher struct {
	ID            int
	workerNum     int
	taskGenerator TaskGenerator
	checker       Checker
	backlogNum    int

	taskHandler TaskHandler
	workChecker WorkerChecker

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
func (i *Dispatcher) SetupHandler(handler TaskGenerator){
	i.taskGenerator = handler
}


func (i *Dispatcher) RegistryWorker(handler TaskHandler){
	i.taskHandler = handler
}
func (i *Dispatcher) RegistryChecker(checker Checker){
	i.checker = checker
}

func (i *Dispatcher) RegistryWorkerChecker(checker WorkerChecker){
	i.workChecker = checker
}


func (i *Dispatcher) Workers(consumerNum int){
	i.workerNum = consumerNum
}

func (i *Dispatcher)Dispatch(wg *sync.WaitGroup){
	defer func(){
		if err := recover(); nil != err {
			i.checker(err)
			logger.Warn("Dispatcher occurs an error, respawn again, err:", err)
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
		worker.handler = i.taskHandler
		worker.workChecker = i.workChecker
		workers = append(workers, worker)
		go func(worker *Worker) {
			defer wg.Done()
			defer func() {
				if err := recover(); nil != err {
					i.workChecker(err)
					//logger.Err("worker occurs an error", err)
				}
			}()
			for{
				select {
				case task,_ = <-pipeline:
					worker.work(task)
					if gloableSingal == GracefulExit {
						return
					}
				}
			}
		}(worker)
	}
}

