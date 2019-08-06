package go_log_test

import (
    "testing"
    log "github.com/weizhouBlue/go_log"
    "os"
)


func Test_stdout(t *testing.T){
    log.Config(  log.Info , "test module" , "" ) 

    log.Log( log.Debug , "this is an debug message" )
    log.Log( log.Info ,  "this is an info message" )
    log.Log( log.Err ,  "this is an err message" )
    //log.Log( log.Panic , "this is an panic message" )
}


func Test_file(t *testing.T){
    fileName := "./log"
    log.Config(  log.Info , "test module" , fileName ) 

    log.Log( log.Debug , "this is an debug message" )
    log.Log( log.Info ,  "this is an info message" )
    log.Log( log.Err ,  "this is an err message" )

    log.Close()
    os.Remove(fileName)

}


