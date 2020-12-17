package go_log

import (
    "log"
    "fmt"
    "os"
    "runtime"
    "strconv"
    "strings"
    "time"
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
    prefix string
}

var recorder Clog

func toString(a interface{}) string {
    if v, p := a.(int); p {
        return strconv.Itoa(v)
    }
    if v, p := a.(int16); p {
        return strconv.Itoa(int(v))
    }
    if v, p := a.(int32); p {
        return strconv.Itoa(int(v))
    }
    if v, p := a.(uint); p {
        return strconv.Itoa(int(v))
    }
    if v, p := a.(float32); p {
        return strconv.FormatFloat(float64(v), 'f', -1, 32)
    }
    if v, p := a.(float64); p {
        return strconv.FormatFloat(v, 'f', -1, 32)
    }
    return ""
}

func existFile( filePath string ) bool {
	   if info , err := os.Stat(filePath) ; err==nil {
		if ! info.IsDir() {
			return true
		}
    }
    return false
}

func getFileName( path string ) string {
    b:=strings.LastIndex(path,"/")
    if b>=0 {
        return path[b+1:]
    }else{
        return path
    }
}

func Config(  level int , prefix , filePath string )  {
	if recorder.file != nil {
		recorder.file.Close()
	}

    recorder.Outputlevel=level
	if len(filePath) > 0 {
		file, err := os.OpenFile( filePath , os.O_CREATE | os.O_WRONLY | os.O_APPEND , 0755) 
		if err!=nil {
			panic(err)
		}
		recorder.file=file
    	recorder.Logger=log.New( file , "["+prefix+"] " , log.Ltime )
        Log(Debug , "initialize clog , output to file %s \n" , filePath  )
    	//fmt.Printf(  "initialize clog , output to file %s \n" , filePath )
	}else{
    	recorder.Logger=log.New( os.Stdout , "["+prefix+"] "  , log.Ltime  )
        Log(Debug , "initialize clog , output to stdout \n"  )    	
        //fmt.Printf("initialize clog , output to stdout \n"  )
	}
    recorder.prefix="["+prefix+"] "

    Log(Debug , "initialize clog , set level to %s \n" , log_level[level]  )
    //fmt.Printf("initialize clog , set level to %s \n" , log_level[level] )

}

func Log( level int , format string , v ... interface{}){
	if recorder.Logger == nil {
		fmt.Println("clog module is initialized by default")
		Config( Info , "" , "" )
	}

	if level >= recorder.Outputlevel {

		prefix := "[" + log_level[level] + "] "
	    funcName,filepath ,line,ok := runtime.Caller(1)
        suffix:=""
	    if ok {
	    	file:=getFileName(filepath)
	    	funcname:=getFileName(runtime.FuncForPC(funcName).Name())
	    	suffix = "    [" + file + " " + funcname + " " + toString(line) +  "]"
	    }

        format=strings.TrimRight(format," \n")

        if level >= Err {
            s:=fmt.Sprintf( recorder.prefix + "%v " + prefix + format + suffix + "\n" , time.Now().Format("15:04:05") )
            fmt.Fprintf( os.Stderr ,  s  ,  v... ) 
            if recorder.file!=nil{
                // record to file
                recorder.Logger.Printf( prefix + format + suffix + "\n" , v... )
            }
        }else{
            recorder.Logger.Printf( prefix + format + suffix + "\n" , v... )
        }
	}


	if level == Panic {
		panic( "error! PANIC happen" )
	}
}

func Close(){
	if recorder.file !=nil {
		recorder.file.Close()
	}
}



