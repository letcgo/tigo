package tigo

type TaskHandler func(*Task)(error)
type WorkerChecker func(interface{})
type IWorker interface {
	work(task *Task)error
}

type Worker struct {
	IWorker
	ID          int
	handler     TaskHandler
	workChecker WorkerChecker
}


func (i *Worker) work(task *Task)error{
	i.handler(task)
	return nil
}

