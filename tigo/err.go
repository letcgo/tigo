package tigo

func CheckErr(e error){
	if nil != e {
		if nil != logger {
			logger.Err(e)
		}else{
			panic(e)
		}
	}
}
