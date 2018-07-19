package app

import (
	"log"
	"io"
)

var logger ILogger

type ILogger interface {
	Info(v ...interface{})
	Notice(v ...interface{})
	Warn(v ...interface{})
	Err(v ...interface{})
}

type Logger struct {
	logger *log.Logger
}

func (i *Logger)Info(v ...interface{}){
	v = append([]interface{}{"[INFO]"}, v...)
	i.logger.Println(v...)
}

func (i *Logger)Notice(v ...interface{}){
	v = append([]interface{}{"[NOTICE]"}, v...)
	i.logger.Println(v...)
}

func (i *Logger)Warn(v ...interface{}){
	v = append([]interface{}{"[WARN]"}, v...)
	i.logger.Println(v...)
}

func (i *Logger)Err(v ...interface{}){
	v = append([]interface{}{"[ERR]"}, v...)
	i.logger.Println(v...)
}

func NewLogger(out io.Writer, prefix string, flag int)ILogger{
	return  &Logger{logger: log.New(out, prefix, flag)}
}

func RegistryLogger(log ILogger){
	logger = log
}

func GetLogger()ILogger{
	return logger
}