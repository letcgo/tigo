package app

import "log"

var logger *Logger

type ILogger interface {
	Info()
	Notice()
	Warn()
	Err()
}

type Logger struct {
	ILogger
	log.Logger
}

func (i *Logger)Info(v ...interface{}){
	v = append(v, "INFO")
	i.Println(v...)
}

func (i *Logger)Notice(v ...interface{}){
	v = append(v, "NOTICE")
	i.Println(v...)
}

func (i *Logger)Warn(v ...interface{}){
	v = append(v, "WARN")
	i.Println(v...)
}

func (i *Logger)Err(v ...interface{}){
	v = append(v, "ERR")
	i.Println(v...)
}

func NewLogger()*Logger{
	return &Logger{}
}

func RegistryLogger(log *Logger){
	logger = log
}