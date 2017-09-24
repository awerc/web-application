package xml_validator

import (
	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"

	"errors"
	"io/ioutil"
	"os"
	"unsafe"
)

func Validate(xmlfilename string, xsdfilename string) error {

	xsdFile, err := os.Open(xsdfilename)
	if err != nil {
		return err
	}
	defer xsdFile.Close()

	byteValue, err := ioutil.ReadAll(xsdFile)
	if err != nil {
		return err
	}

	xsdSchema, err := xsd.ParseSchema(byteValue)
	if err != nil {
		return err
	}

	doc := golibxml.ParseFile(xmlfilename)
	if doc == nil {
		return errors.New("Error parsing document")
	}
	defer doc.Free()

	err = xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr)))
	if err != nil {
		return err
	}

	return nil
}
