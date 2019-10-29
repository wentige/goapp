package dh

import (
	"encoding/xml"
)

// Request

type Login struct {
	UserId   string `xml:"USERID"`
	Password string `xml:"PASSWORD"`
}

type XmlRequest struct {
	XMLName  xml.Name `xml:"XMLFORMPOST"`
	Request  string   `xml:"REQUEST"`
	Login    Login    `xml:"LOGIN"`
	Partnums []string `xml:"PARTNUM"`
}

// Response

type BranchQty struct {
	Branch      string `xml:"BRANCH"`
	Qty         string `xml:"QTY"`
	InStockDate string `xml:"INSTOCKDATE"`
}

type Item struct {
	PartNum   string      `xml:"PARTNUM"`
	UnitPrice string      `xml:"UNITPRICE"`
	BranchQty []BranchQty `xml:"BRANCHQTY"`
	TotalQty  string      `xml:"TOTALQTY"`
	Message   string      `xml:"MESSAGE"`
}

type XmlResponse struct {
	XMLName xml.Name `xml:"XMLRESPONSE"`
	Items   []Item   `xml:"ITEM"`
	Status  string   `xml:"STATUS"`
}
