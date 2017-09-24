package xml_parser

import (
	"encoding/xml"
)

type Configuration struct {
	XMLName     xml.Name    `xml:"configuration"`
	AppSettings AppSettings `xml:"appSettings"`
	Http        Http        `xml:"http"`
	Database    []Database  `xml:"database"`
}

type AppSettings struct {
	XMLName xml.Name `xml:"appSettings"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
}

type Http struct {
	XMLName xml.Name `xml:"http"`
	Listen  int      `xml:"listen"`
	Timeout int      `xml:"timeout"`
}

type Database struct {
	XMLName xml.Name `xml:"database"`
	Host    string   `xml:"host"`
	Port    int      `xml:"port"`
	User    string   `xml:"user"`
	Pass    string   `xml:"pass"`
}
