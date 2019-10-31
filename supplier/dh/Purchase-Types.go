package dh

import "encoding/xml"

type XmlOrderEntry struct {
	XMLName xml.Name `xml:"XMLFORMPOST"`
	Request string   `xml:"REQUEST"`
	Login   struct {
		UserId   string `xml:"USERID"`
		Password string `xml:"PASSWORD"`
	} `xml:"LOGIN"`
	Header struct {
		OnlyBranch       string `xml:"ONLYBRANCH"`
		Branches         string `xml:"BRANCHES"`
		PartshipAllow    string `xml:"PARTSHIPALLOW"`
		BackorderAllow   string `xml:"BACKORDERALLOW"`
		DropshipPW       string `xml:"DROPSHIPPW"`
		ShipToName       string `xml:"SHIPTONAME"`
		ShipToAttn       string `xml:"SHIPTOATTN"`
		ShipToAddress    string `xml:"SHIPTOADDRESS"`
		ShipToAddress2   string `xml:"SHIPTOADDRESS2"`
		ShipToCity       string `xml:"SHIPTOCITY"`
		ShipToProvince   string `xml:"SHIPTOPROVINCE"`
		ShipToPostalCode string `xml:"SHIPTOPOSTALCODE"`
		ShipCarrier      string `xml:"SHIPCARRIER"`
		ShipService      string `xml:"SHIPSERVICE"`
		PoNum            string `xml:"PONUM"`
		Remarks          string `xml:"REMARKS"`
	} `xml:"ORDERHEADER"`
	Items struct {
		Item struct {
			Partnum string `xml:"PARTNUM"`
			Qty     string `xml:"QTY"`
			Branch  string `xml:"BRANCH"`
		} `xml:"ITEM"`
	} `xml:"ORDERITEMS"`
}

type XmlOrderResponse struct {
	XMLName  xml.Name `xml:"XMLRESPONSE"`
	OrderNum string   `xml:"ORDERNUM"`
	Status   string   `xml:"STATUS"`
	Message  string   `xml:"MESSAGE"`
}
