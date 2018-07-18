package app


type IMonitor interface {
	Notify(v ...interface{})
}


type Monitor struct {
}

func (i *Monitor)Notify(v ...interface{}){
	//todo sms, mail, phone, wechat
	println("Notify", v)
}

var monitor IMonitor

func NewMonitor()*Monitor{
	return new(Monitor)
}

func RegistryMonitor(m IMonitor){
	monitor = m
}