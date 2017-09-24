package xml_parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

func Parse(filename string) (configuration Configuration, err error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return configuration, err
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return configuration, err
	}
	err = xml.Unmarshal(byteValue, &configuration)

	return configuration, err
}
