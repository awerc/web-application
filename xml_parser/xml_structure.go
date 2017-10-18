//Structure of xml file
package xml_parser

import (
	"encoding/xml"
)

type Configuration struct {
	XMLName xml.Name `xml:"configuration"`
	Connect Connect  `xml:"connect"`
	User    User     `xml:"user"`
}

type Connect struct {
	XMLName       xml.Name `xml:"connect"`
	Host          string   `xml:"host"`
	Port          int      `xml:"port"`
	Db            string   `xml:"db"`
	Authorization string   `xml:"authorization"`
	Timeout       int      `xml:"timeout"`
}

type User struct {
	XMLName  xml.Name `xml:"user"`
	Type     string   `xml:"type"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
}
