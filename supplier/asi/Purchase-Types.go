package asi

import "encoding/xml"

// request

type ASIOrderRequest struct {
	XMLName xml.Name `xml:"ASIOrderRequest"`
	Custid  string   `xml:"custid,attr"`
	Custpo  string   `xml:"custpo,attr"`
	Cert    string   `xml:"cert,attr"`
	Header  struct {
		Shipto struct {
			Name     string `xml:"name"`
			Address1 string `xml:"address1"`
			Address2 string `xml:"address2"`
			Address3 string `xml:"address3"`
			City     string `xml:"city"`
			State    string `xml:"state"`
			Zip      string `xml:"zip"`
			Country  string `xml:"country"`
		} `xml:"shipto"`
		Shipinfo struct {
			Via          string `xml:"via"`
			Instructions string `xml:"instructions"`
		} `xml:"shipinfo"`
	} `xml:"header"`
	Detail struct {
		Line struct {
			Sku    string `xml:"sku"`
			Qty    string `xml:"qty"`
			Price  string `xml:"price"`
			Branch string `xml:"branch"`
		} `xml:"line"`
	} `xml:"detail"`
}

// response

type ASIOrderReply struct {
	XMLName xml.Name `xml:"ASIOrderReply"`
	Order   struct {
		Orderid string `xml:"orderid"`
	} `xml:"order"`
}
