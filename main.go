package main

import (
	"flag"
	"fmt"
	"os"
	"web-application/logger"
	"web-application/xml_parser"
	"web-application/xml_validator"
)

var (
	configname  string
	xsdFilename string
	log         string
)

func init() {
	flag.StringVar(&configname, "config", "config.xml", "name of config file")
	flag.StringVar(&xsdFilename, "xsd", "xsd", "name of xsd file")
	flag.StringVar(&log, "log", "log.txt", "name of log file")

	flag.Parse()
}

func main() {
	var logger = logger.Create(nil, logger.Debug|logger.Info, logger.DateTime|logger.Shortfile)
	defer logger.Close()
	logger.AddOutput(os.Stdout)
	logger.AddOutputFile(log)

	err := xml_validator.Validate(configname, xsdFilename)
	if err != nil {
		logger.Info(err)
	}

	configuration, err := xml_parser.Parse(configname)
	if err != nil {
		logger.Info(err)
	}

	fmt.Printf("%v\n", configuration.Database[0].User)
}
