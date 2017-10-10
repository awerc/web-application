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
	configFilename string
	xsdFilename    string
	logFilename    string
)

func init() {
	flag.StringVar(&configFilename, "config", "config.xml", "name of config file")
	flag.StringVar(&xsdFilename, "xsd", "xsd", "name of xsd file")
	flag.StringVar(&logFilename, "log", "log.txt", "name of log file")

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
	}

	err = xml_validator.Validate(configuration)
	if err != nil {
		logger.Info(err)
	}

	spew.Dump(configuration)
}
