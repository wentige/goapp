package newegg

import (
	"encoding/xml"
)

type Error struct {
	Code    string `xml:"Code"`
	Message string `xml:"Message"`
}

type Errors struct {
	XMLName xml.Name `xml:"Errors"`
	Xsd     string   `xml:"xsd,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Error   []Error  `xml:"Error"`
}
