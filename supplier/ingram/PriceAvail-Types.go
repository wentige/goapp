package ingram

import (
	"encoding/xml"
)

// Request

type RequestHeader struct {
	SenderID      string `xml:"SenderID"`
	ReceiverID    string `xml:"ReceiverID"`
	CountryCode   string `xml:"CountryCode"`
	LoginID       string `xml:"LoginID"`
	Password      string `xml:"Password"`
	TransactionID string `xml:"TransactionID"`
}

type RequestItem struct {
	SKU string `xml:"SKU,attr"`
	Qty string `xml:"Quantity,attr"`
}

type PNARequest struct {
	XMLName    xml.Name      `xml:"PNARequest"`
	Version    string        `xml:"Version"`
	Header     RequestHeader `xml:"TransactionHeader"`
	Items      []RequestItem `xml:"PNAInformation"`
	ShowDetail string        `xml:"ShowDetail"`
}

// Response

type Branch struct {
	Name         string `xml:"Name,attr"`
	ID           string `xml:"ID,attr"`
	Availability string `xml:"Availability"`
	OnOrder      string `xml:"OnOrder"`
	ETADate      string `xml:"ETADate"`
}

type Item struct {
	SKU                string   `xml:"SKU,attr"`
	Qty                string   `xml:"Quantity,attr"`
	Price              string   `xml:"Price"`
	SpecialPriceFlag   string   `xml:"SpecialPriceFlag"`
	PartNumber         string   `xml:"ManufacturerPartNumber"`
	PartNumberOccurs   string   `xml:"ManufacturerPartNumberOccurs"`
	VendorNumber       string   `xml:"VendorNumber"`
	Description        string   `xml:"Description"`
	ReserveFlag        string   `xml:"ReserveInventoryFlag"`
	AvailableRebQty    string   `xml:"AvailableRebQty"`
	Branches           []Branch `xml:"Branch"`
	UPC                string   `xml:"UPC"`
	CustomerPartNumber string   `xml:"CustomerPartNumber"`
}

type ErrorStatus struct {
	ErrorNumber string `xml:"ErrorNumber,attr"`
}

type ResponseHeader struct {
	SenderID      string      `xml:"SenderID"`
	ReceiverID    string      `xml:"ReceiverID"`
	ErrorStatus   ErrorStatus `xml:"ErrorStatus"`
	DocumentID    string      `xml:"DocumentID"`
	TransactionID string      `xml:"TransactionID"`
	TimeStamp     string      `xml:"TimeStamp"`
}

type PNAResponse struct {
	XMLName xml.Name       `xml:"PNAResponse"`
	Version string         `xml:"Version"`
	Header  ResponseHeader `xml:"TransactionHeader"`
	Items   []Item         `xml:"PriceAndAvailability"`
}
