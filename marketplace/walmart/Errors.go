package walmart

import (
	"encoding/xml"
)

type Errors struct {
	XMLName xml.Name `xml:"errors"`
	Ns2     string   `xml:"ns2,attr"`
	Ns3     string   `xml:"ns3,attr"`
	Ns4     string   `xml:"ns4,attr"`
	Error   struct {
		Code        string `xml:"code"`
		Field       string `xml:"field"`
		Description string `xml:"description"`
		Info        string `xml:"info"`
		Severity    string `xml:"severity"`
		Category    string `xml:"category"`
		Causes      string `xml:"causes"`
		Identifiers string `xml:"errorIdentifiers"`
	} `xml:"error"`
}
