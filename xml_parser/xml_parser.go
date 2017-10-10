//Package that parse xml file into custom structure
//described in xml_structure.go
package xml_parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

//A function that parse filename file into Configuration structure
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
