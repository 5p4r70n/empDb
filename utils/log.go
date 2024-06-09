package utils

import (
		"github.com/zerodha/logf"
		"time"
		"os"
	)



func Log() (func()*logf.Logger) {
	
	file,err:=os.OpenFile("./log.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0777);
	if err!=nil{panic("Unable to create log file")}
	
	return func() *logf.Logger{

		logger := logf.New(logf.Opts{
			Writer:               file,
			EnableColor:          false,
			Level:                logf.DebugLevel,
			CallerSkipFrameCount: 3,
			EnableCaller:         true,
			TimestampFormat:      time.RFC3339Nano,
		})
		
		return &logger
	}
}


