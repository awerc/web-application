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
		AppSettings: AppSettings{
			XMLName: xml.Name{
				Space: "",
				Local: "appSettings",
			},
			Type: "123",
			Name: "UserName",
		},
		Http: Http{
			XMLName: xml.Name{
				Space: "",
				Local: "http",
			},
			Listen:  8080,
			Timeout: 5000,
		},
		Database: Database{
			XMLName: xml.Name{
				Space: "",
				Local: "database",
			},
			Name: "postgres",
			Host: "localhost",
			User: "admin",
			Pass: "password",
		},
	}
)

func TestParsing(t *testing.T) {

	actual, _ := Parse("../config.xml")
	if actual != expected {
		t.Errorf("Test failed, expected: '%v',\ngot:  '%v'", expected, actual)
	}
}
