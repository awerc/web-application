package main

import (
	"flag"
	"os"
	"web-application/logger"
	"web-application/xml_parser"
	"web-application/xml_validator"

	"github.com/davecgh/go-spew/spew"
)

var (
	configFilename string //String that contains path to config file
	logFilename    string //String that contains path to log file
)

//Reading command line parameters before programm execution
func init() {
	flag.StringVar(&configFilename, "config", "config.xml", "name of config file")
	//Command line parameter for config file. Default is config.xml
	flag.StringVar(&logFilename, "log", "log.txt", "name of log file")
	//Command line parameter for log file. Default is log.txt

	flag.Parse()
}

func main() {
	var logger = logger.Initialize("[WEB APP]", logger.Debug|logger.Info, logger.DateTime|logger.Shortfile)
	defer logger.Close()
	logger.AddOutput(os.Stdout)
	logger.AddOutputFile(logFilename)

	configuration, err := xml_parser.Parse(configFilename)
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}

	err = xml_validator.Validate(configuration)
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}

	spew.Dump(configuration)
}
