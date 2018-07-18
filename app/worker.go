package app

type TaskHandler func(*Task)(error)
type WorkChecker func(error)
type IWorker interface {
	work(task *Task)error
}

type Worker struct {
	IWorker
	ID      int
	handler TaskHandler
	workChecker WorkChecker
}


func (i *Worker) work(task *Task, )error{
	if nil != i.workChecker {
		i.workChecker(i.handler(task))
	}else{
		i.handler(task)
	}
	return nil
}

