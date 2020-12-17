package go_log_test

import (
    "testing"
    log "github.com/weizhouBlue/go_log"
    "os"
    "path/filepath"
)


func Test_stdout(t *testing.T){
    var BinName=filepath.Base(os.Args[0])

    log.Config(  log.Info , BinName , "" ) 


    log.Log( log.Debug , "    this is an debug message %v \n", "1111" )
    log.Log( log.Debug , "this is an debug message \n" )
    log.Log( log.Info ,  "             this is an info message  %v \n  \n" , 123 )
    log.Log( log.Info ,  "this is an info message \n" )
    log.Log( log.Err ,  "this is an err message  %v \n" , 23432 )
    log.Log( log.Err ,  "this is an err message \n" )
    //log.Log( log.Panic , "this is an panic message" )
}


func Test_file(t *testing.T){
    var BinName=filepath.Base(os.Args[0])

    fileName := "./log"
    log.Config(  log.Info , BinName , fileName ) 

    log.Log( log.Debug , "this is an debug message \n" )
    log.Log( log.Info ,  "this is an info message \n" )
    log.Log( log.Err ,  "this is an err message \n" )

    log.Close()
    //os.Remove(fileName)

}



