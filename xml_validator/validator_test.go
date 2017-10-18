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
}

func GetConfigWithWrongPort() Configuration {
	return Configuration{
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
			Port:          5,
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
}

func GetConfigWithWrongTimeout() Configuration {
	return Configuration{
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
			Timeout:       1,
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
}

func GetConfigWithWrongUsername() Configuration {
	return Configuration{
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
			Login:    "*/-",
			Password: "pass",
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
