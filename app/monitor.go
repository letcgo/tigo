package app


type IMonitor interface {
	Notify(v ...interface{})
}


type Monitor struct {
}

func (i *Monitor)Notify(v ...interface{}){
	//todo sms, mail, phone, wechat
	//logger.Info(v...)
}

var monitor IMonitor

func NewMonitor()*Monitor{
	return new(Monitor)
}
