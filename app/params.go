package app

import "flag"

func init()  {
	var help bool
	flag.BoolVar(&help, "help", false, "man")
	flag.Parse()
	if help {
		flag.Usage()
	}
}

