package synnex

import "encoding/xml"

type OrderRequest struct {
	XMLName    xml.Name `xml:"SynnexB2B"`
	Credential struct {
		UserID   string `xml:"UserID"`
		Password string `xml:"Password"`
	} `xml:"Credential"`
	Detail struct {
		CustomerNumber string `xml:"CustomerNumber"`
		PONumber       string `xml:"PONumber"`
		DropShipFlag   string `xml:"DropShipFlag"`
		Shipment       struct {
			ShipFromWarehouse string `xml:"ShipFromWarehouse"`
			ShipTo            struct {
				AddressName1 string `xml:"AddressName1"`
				AddressLine1 string `xml:"AddressLine1"`
				AddressLine2 string `xml:"AddressLine2"`
				City         string `xml:"City"`
				State        string `xml:"State"`
				ZipCode      string `xml:"ZipCode"`
				Country      string `xml:"Country"`
			} `xml:"ShipTo"`
			ShipToContact struct {
				ContactName  string `xml:"ContactName"`
				PhoneNumber  string `xml:"PhoneNumber"`
				EmailAddress string `xml:"EmailAddress"`
			} `xml:"ShipToContact"`
			ShipMethod struct {
				Code string `xml:"Code"`
			} `xml:"ShipMethod"`
		} `xml:"Shipment"`
		Items struct {
			Item struct {
				LineNumber string `xml:"lineNumber,attr"`
				SKU        string `xml:"SKU"`
				UnitPrice  string `xml:"UnitPrice"`
				Quantity   string `xml:"OrderQuantity"`
			} `xml:"Item"`
		} `xml:"Items"`
	} `xml:"OrderRequest"`
}

type OrderResponse struct {
	XMLName xml.Name `xml:"SynnexB2B"`
	Detail  struct {
		CustomerNumber      string `xml:"CustomerNumber"`
		PONumber            string `xml:"PONumber"`
		Code                string `xml:"Code"`
		ResponseDateTime    string `xml:"ResponseDateTime"`
		ResponseElapsedTime string `xml:"ResponseElapsedTime"`
		Items               struct {
			Item struct {
				LineNumber        string `xml:"lineNumber,attr"`
				SKU               string `xml:"SKU"`
				Quantity          string `xml:"OrderQuantity"`
				Code              string `xml:"Code"`
				OrderNumber       string `xml:"OrderNumber"`
				OrderType         string `xml:"OrderType"`
				ShipFromWarehouse string `xml:"ShipFromWarehouse"`
				SynnexInternalRef string `xml:"SynnexInternalReference"`
			} `xml:"Item"`
		} `xml:"Items"`
	} `xml:"OrderResponse"`
}
