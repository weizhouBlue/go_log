package go_log_test

import (
	log "github.com/weizhouBlue/go_log"
	"os"
	"path/filepath"
	"testing"
)

func Test_stdout(t *testing.T) {
	var BinName = filepath.Base(os.Args[0])

	log.Config(log.Info, BinName, os.Stdout )

	log.Log(log.Debug, "    this is an debug message %v \n", "1111")
	log.Log(log.Debug, "this is an debug message \n")
	log.Log(log.Info, "             this is an info message  %v \n  \n", 123)
	log.Log(log.Info, "this is an info message \n")
	log.Log(log.Err, "this is an err message  %v \n", 23432)
	log.Log(log.Err, "this is an err message \n")
	//log.Log( log.Panic , "this is an panic message" )
}

func Test_file(t *testing.T) {
	var BinName = filepath.Base(os.Args[0])

	fileName := "./log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	log.Config(log.Info, BinName, file)

	log.Log(log.Debug, "this is an debug message \n")
	log.Log(log.Info, "this is an info message \n")
	log.Log(log.Err, "this is an err message \n")


	//os.Remove(fileName)

}
