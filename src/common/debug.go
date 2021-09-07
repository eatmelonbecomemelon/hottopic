package common



import (
	"fmt"
	"log"

	"runtime"
)

var (
	ShowDebug   = true
	ShowWarning = true
	ShowError   = true
	ShowInfo    = true
)
const sysflag = "hottopic"

var FULL_LOG = true
var START_LOG = false
var LOG_FILE_PATH = "log"
var rootDir = "./"

func SetLogLevel(logLevel string) {
	switch logLevel {
	case "DEBUG":
		ShowDebug, ShowInfo, ShowWarning, ShowError = true, true, true, true
	case "INFO":
		ShowDebug, ShowInfo, ShowWarning, ShowError = false, true, true, true
	case "WARNING":
		ShowDebug, ShowInfo, ShowWarning, ShowError = false, false, true, true
	case "ERROR":
		ShowDebug, ShowInfo, ShowWarning, ShowError = false, false, false, true
	}
}

func output(info string, format string, v ...interface{}) {
	if FULL_LOG {
		pc, file, line, _ := runtime.Caller(2)
		short := file

		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		f := runtime.FuncForPC(pc)
		fn := f.Name()

		for i := len(fn) - 1; i > 0; i-- {
			if fn[i] == '.' {
				fn = fn[i+1:]
				break
			}
		}

		if format == "" {
			log.Printf("|%v|%v|%v|%v()|%v|%v", info, sysflag, short, fn, line, fmt.Sprintln(v...))
		} else {
			log.Printf("|%v|%v|%v|%v()|%v|%v", info, sysflag, short, fn, line, fmt.Sprintf(format, v...))
		}

	} else {
		if format == "" {
			log.Printf("[%s]|%v", info, fmt.Sprintln(v...))
		} else {
			log.Printf("[%s]|%v", info, fmt.Sprintf(format, v...))
		}
	}

}

func Debug(v ...interface{}) {
	if ShowDebug {
		output("DEBUG", "", v...)
	}
}
func Debugf(format string, v ...interface{}) {
	if ShowDebug {
		output("DEBUG", format, v...)
	}
}

func Error(v ...interface{}) {
	if ShowError {
		output("ERROR", "", v...)
	}
}
func Errorf(format string, v ...interface{}) {
	if ShowError {
		output("ERROR", format, v...)
	}
}

func Warning(v ...interface{}) {
	if ShowWarning {
		output("WARNING", "", v...)
	}
}
func Warningf(format string, v ...interface{}) {
	if ShowWarning {
		output("WARNING", format, v...)
	}
}

func Info(v ...interface{}) {
	if ShowInfo {
		output("INFO", "", v...)
	}
}
func Infof(format string, v ...interface{}) {
	if ShowInfo {
		output("INFO", format, v...)
	}
}

