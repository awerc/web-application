package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	Debug = 1 << iota
	Info
)

const (
	DateTime  = 1 << iota // the Date and Time in the local time zone: 2009/01/23 - 15:04:05
	Longfile              // full file name and line number: /a/b/c/d.go:23
	Shortfile             // final file name element and line number: d.go:23 overrides Longfile
)

type Logger struct {
	out        []*os.File //Destinations for output
	prefix     string     //Smth that prints before message
	logLevel   int        //Debug | Info
	flag       int        //DateTime | Longfile | Shortfile
	timeFormat string     //defaul is "2006/01/02 - 15:04:05"
}

//After Initializing u have to add outputs with
//AddOutputFile or AddOutput.
//
//Available logLevels: Debug | Info.
//Available flags: DateTime | Longfile | Shortfile.
//
//Use 'defer logger.Close()' after creating an instance
//to close all opened files
func Initialize(prefix string, logLevels int, flag int) *Logger {
	return &Logger{out: []*os.File{}, logLevel: logLevels, flag: flag, timeFormat: "2006/01/02 - 15:04:05", prefix: prefix}
}

//Adds an output file to the list using the file name
func (logger *Logger) AddOutputFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	}
	logger.out = append(logger.out, file)
	return nil
}

//Adds an output to the list
func (logger *Logger) AddOutput(out *os.File) error {
	if out != nil {
		logger.out = append(logger.out, out)
	} else {
		return errors.New("out is equal nil")
	}
	return nil
}

//Delete all outputs
func (logger *Logger) ClearOutputs() error {
	logger.out = []*os.File{}
	return nil
}

//Print Debug message
func (logger *Logger) Debug(v ...interface{}) {
	if logger.logLevel&Debug != 0 {
		for _, element := range logger.out {
			str := logger.createLogString(v, "DEBUG")
			if element == os.Stdout && runtime.GOOS != "windows" {
				element.WriteString("\x1b[34m\x1b[1m" + str + "\x1b[0m")
			} else {
				element.WriteString(str)
			}
		}
	}
}

//Print Info message
func (logger *Logger) Info(v ...interface{}) {
	if logger.logLevel&Info != 0 {
		for _, element := range logger.out {
			str := logger.createLogString(v, "INFO")
			if element == os.Stdout && runtime.GOOS != "windows" {
				element.WriteString("\x1b[32m\x1b[1m" + str + "\x1b[0m")
			} else {
				element.WriteString(str)
			}
		}
	}
}

// ChangeLevel change the level of logging.
// Available levels "INFO" | "DEBUG"
func (logger *Logger) ChangeLevel(flag int) error {
	if flag > 3 || flag < 1 {
		return errors.New("Wrong flags to be set (possible: Debug and Info")
	}
	logger.logLevel = flag
	return nil
}

// SetTimeFormat sets string format for time.Time.Format() method
// Default is "2006/01/02 - 15:04:05"
func (logger *Logger) SetTimeFormat(format string) {
	logger.timeFormat = format
}

//Method for obtaining a file and a call string for logging
func (logger *Logger) getFileAndLine() (string, int) {
	_, file, line, _ := runtime.Caller(3)
	if logger.flag&Shortfile != 0 {
		file = filepath.Base(file)
	}
	return file, line
}

//Method for creating output message
//level can be "DEBUG" or "INFO"
//The debug displays detailed information about the given parameters
func (logger *Logger) createLogString(v []interface{}, level string) string {
	file, line := logger.getFileAndLine()
	var out string
	if logger.flag&DateTime != 0 {
		now := time.Now().Format(logger.timeFormat)
		out += fmt.Sprint(logger.prefix, " [", level, "] ", now, "  ", file, ":", line, "  ▶  ")
	} else {
		out += fmt.Sprint(logger.prefix, " [", level, "] ", file, ":", line, "  ▶  ")
	}
	for i, value := range v {
		if level == "DEBUG" {
			value = spew.Sdump(value)
		}

		out += fmt.Sprintf("%+v", value)
		if level != "DEBUG" && i < len(v)-1 {
			out += fmt.Sprint(" | ")
		}
	}
	out += "\n"
	return out
}

//Close all outputs
func (logger *Logger) Close() error {
	for _, element := range logger.out {
		err := element.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
