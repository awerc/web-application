package xml_parser

import (
	"encoding/xml"
	"testing"
)

var (
	expected = Configuration{
		XMLName: xml.Name{
			Space: "",
			Local: "configuration",
		},
		Connect: Connect{
			XMLName: xml.Name{
				Space: "",
				Local: "connect",
			},
			Host:          "localhost",
			Port:          8080,
			Db:            "postgre",
			Authorization: "yes",
			Timeout:       1000,
		},
		User: User{
			XMLName: xml.Name{
				Space: "",
				Local: "user",
			},
			Type:     "admin",
			Login:    "adm",
			Password: "pass",
		},
	}
)

func TestParsing(t *testing.T) {
	actual, _ := Parse("../config.xml")
	if actual != expected {
		t.Errorf("Test failed, expected: '%v',\ngot:  '%v'", expected, actual)
	}
}
