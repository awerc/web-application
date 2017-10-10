package xml_validator

import (
	"encoding/xml"
	"testing"
	. "web-application/xml_parser"
)

func GetValidConfig() Configuration {
	return Configuration{
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
}

func GetConfigWithWrongPort() Configuration {
	return Configuration{
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
			Listen:  100,
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
}

func GetConfigWithWrongTimeout() Configuration {
	return Configuration{
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
			Timeout: 1,
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
}

func GetConfigWithWrongUsername() Configuration {
	return Configuration{
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
			User: "@#$%^",
			Pass: "password",
		},
	}
}

func TestValidConfig(t *testing.T) {
	var expected interface{} = nil
	actual := Validate(GetValidConfig())
	if actual != expected {
		t.Errorf("Test failed, expected: '%v',\ngot:  '%v'", expected, actual)
	}
}

func TestWrongPort(t *testing.T) {
	expected := "Wrong port"
	actual := Validate(GetConfigWithWrongPort())
	if actual.Error() != expected {
		t.Errorf("Test failed, expected: '%v',\ngot:  '%v'", expected, actual)
	}
}

func TestWrongTimeout(t *testing.T) {
	expected := "Too short timeout"
	actual := Validate(GetConfigWithWrongTimeout())
	if actual.Error() != expected {
		t.Errorf("Test failed, expected: '%v',\ngot:  '%v'", expected, actual)
	}
}

func TestWrongUsername(t *testing.T) {
	expected := "Wrong user name"
	actual := Validate(GetConfigWithWrongUsername())
	if actual.Error() != expected {
		t.Errorf("Test failed, expected: '%v',\ngot:  '%v'", expected, actual)
	}
}
