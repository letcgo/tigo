package tigo

import (
	"flag"
	"os"
)
var help bool
var envPath string
func init()  {
	flag.BoolVar(&help, "help", false, `help info`)
	flag.StringVar(&envPath, "e", "", "<env file path>, default: ./.env")
	flag.Parse()
	if help || len(envPath)==0 {
		flag.Usage()
		os.Exit(1)
	}
}

func EnvPath()string  {
	return envPath
}
