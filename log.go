package go_log

import (
    "log"
    "fmt"
    "os"
)

const (
    Debug=iota
    Info
    Err
    Panic
)
var log_level = [...]string {
	"debug" ,
	"info" ,
	"error" ,
	"Panic" ,
}

type Clog struct {
    Logger *log.Logger
    Outputlevel int 
    file *os.File
}

var recorder Clog

func existFile( filePath string ) bool {
	   if info , err := os.Stat(filePath) ; err==nil {
		if ! info.IsDir() {
			return true
		}
    }
    return false
}

func Conf(  level int , prefix , filePath string )  {
	if recorder.file != nil {
		recorder.file.Close()
	}

	if len(filePath) > 0 {
		file, err := os.OpenFile( filePath , os.O_CREATE | os.O_WRONLY | os.O_APPEND , 0755) 
		if err!=nil {
			panic(err)
		}
		recorder.file=file
    	recorder.Logger=log.New( file , "["+prefix+"] " , log.Ltime | log.Lshortfile   )
    	fmt.Printf("initialize clog , output to file %s \n" , filePath )
	}else{
    	recorder.Logger=log.New( os.Stdout , "["+prefix+"]"  , log.Ltime | log.Lshortfile   )
    	fmt.Printf("initialize clog , output to stdout \n"  )
	}
    recorder.Outputlevel=level
    fmt.Printf("initialize clog , set level to %s \n" , log_level[level] )

}

func Log( level int , format string , v ... interface{}){
	if recorder.Logger == nil {
		fmt.Println("clog module is initialized by default")
		Conf( Info , "" , "" )
	}

	if level >= recorder.Outputlevel {
		recorder.Logger.Printf( "[" + log_level[level] + "] " + format, v... )
	}

	if level == Panic {
		panic( "PANIC happen" )
	}
}

func Close(){
	if recorder.file !=nil {
		recorder.file.Close()
	}
}


