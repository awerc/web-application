package main

import (
	"flag"
	"fmt"
	"os"
	"web-application/logger"
	"web-application/xml_parser"
	"web-application/xml_validator"
)

func main() {
	var (
		configname  string
		xsdFilename string
		log         string
	)
	flag.StringVar(&configname, "config", "config.xml", "name of config file")
	flag.StringVar(&xsdFilename, "xsd", "xsd", "name of xsd file")
	flag.StringVar(&log, "log", "log.txt", "name of log file")

	flag.Parse()

	var logger = logger.Create(nil, logger.Debug|logger.Info, logger.DateTime|logger.Shortfile)
	defer logger.Close()
	logger.AddOutput(os.Stdout)
	logger.AddOutputFile(log)
	logger.Debug("wad")

	err := xml_validator.Validate(configname, xsdFilename)
	if err != nil {
		fmt.Println(err)
	}
	configuration, err := xml_parser.Parse(configname)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", configuration.Database[0].User)
}
