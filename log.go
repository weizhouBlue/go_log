package go_log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	Debug = iota
	Info
	Warn
	Err
	Panic
)

var log_level = [...]string{
	"debug",
	"info",
	"warning",
	"error",
	"Panic",
}

type Clog struct {
	Logger      *log.Logger
	Outputlevel int
	prefix      string
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

func getFileName(path string) string {
	b := strings.LastIndex(path, "/")
	if b >= 0 {
		return path[b+1:]
	} else {
		return path
	}
}

func Config(level int, prefix string, out io.Writer) {
	recorder.Outputlevel = level
	recorder.Logger = log.New(out, "["+prefix+"] ", log.Ltime|log.Ldate)
	recorder.prefix = "[" + prefix + "] "
}

func Log(level int, format string, v ...interface{}) {
	if recorder.Logger == nil {
		fmt.Println("clog module is initialized by default")
		Config(Info, "", os.Stdout )
	}

	if level >= recorder.Outputlevel {
		prefix := "[" + log_level[level] + "] "
		funcName, filepath, line, ok := runtime.Caller(1)
		suffix := ""
		if ok {
			file := getFileName(filepath)
			funcname := getFileName(runtime.FuncForPC(funcName).Name())
			suffix = "    [" + file + " " + funcname + " " + toString(line) + "]"
		}

		format = strings.TrimRight(format, " \n")
		recorder.Logger.Printf(prefix+format+suffix+"\n", v...)
	}

	if level == Panic {
		panic("error! PANIC happen")
	}
}


