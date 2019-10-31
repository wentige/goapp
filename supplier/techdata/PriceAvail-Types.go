package techdata

import (
	"encoding/xml"
)

// Request

type RequestHeader struct {
	Username string `xml:"UserName"`
	Password string `xml:"Password"`
	Version  string `xml:"ResponseVersion"`
}

type RequestLineInfo struct {
	RefIDQual string `xml:"RefIDQual"`
	RefID     string `xml:"RefID"`
}

type RequestDetail struct {
	LineInfo []RequestLineInfo `xml:"LineInfo"`
}

type XMLPriceAvailSubmit struct {
	XMLName xml.Name      `xml:"XML_PriceAvailability_Submit"`
	Header  RequestHeader `xml:"Header"`
	Detail  RequestDetail `xml:"Detail"`
}

// Response

type ResponseHeader struct {
	TransSetIDCode  string `xml:"TransSetIDCode"`
	TransControlID  string `xml:"TransControlID"`
	ResponseVersion string `xml:"ResponseVersion"`
}

type Warehouse struct {
	Name string `xml:"IDCode"`
	Code string `xml:"WhseCode"`
	Qty  string `xml:"Qty"`
}

type ErrorInfo struct {
	RefIDQual3 string `xml:"RefIDQual3"`
	RefID3     string `xml:"RefID3"`
	ErrorDesc  string `xml:"ErrorDesc"`
}

type ResponseLineInfo struct {
	RefIDQual1    string      `xml:"RefIDQual1"`
	RefID1        string      `xml:"RefID1"`
	RefIDQual2    string      `xml:"RefIDQual2"`
	RefID2        string      `xml:"RefID2"`
	RefIDQual4    string      `xml:"RefIDQual4"`
	RefID4        string      `xml:"RefID4"`
	ProductDesc   string      `xml:"ProductDesc"`
	PriceIDCode1  string      `xml:"PriceIDCode1"`
	UnitPrice1    string      `xml:"UnitPrice1"`
	PriceIDCode2  string      `xml:"PriceIDCode2"`
	UnitPrice2    string      `xml:"UnitPrice2"`
	EndUserInfo   string      `xml:"RequiredEndUserInfo"`
	LicenseInfo   string      `xml:"RequiredLicenseInfo"`
	ProductWeight string      `xml:"ProductWeight"`
	ItemStatus    string      `xml:"ItemStatus"`
	Warehouses    []Warehouse `xml:"WhseInfo"`
	ErrorInfo     ErrorInfo   `xml:"ErrorInfo"`
}

type ResponseDetail struct {
	LineInfo []ResponseLineInfo `xml:"LineInfo"`
}

type Summary struct {
	NbrOfSegments string `xml:"NbrOfSegments"`
}

type XMLPriceAvailResponse struct {
	XMLName xml.Name       `xml:"XML_PriceAvailability_Response"`
	Header  ResponseHeader `xml:"Header"`
	Detail  ResponseDetail `xml:"Detail"`
	Summary Summary        `xml:"Summary"`
}
