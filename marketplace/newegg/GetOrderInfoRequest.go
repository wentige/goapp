package newegg

import (
	"encoding/xml"
)

type GetOrderInfoRequest struct {
	XMLName       xml.Name `xml:"NeweggAPIRequest"`
	OperationType string   `xml:"OperationType"`
	RequestBody   struct {
		PageIndex       string `xml:"PageIndex"`
		PageSize        string `xml:"PageSize"`
		RequestCriteria struct {
			OrderNumberList struct {
				OrderNumber []string `xml:"OrderNumber"`
			} `xml:"OrderNumberList"`
		} `xml:"RequestCriteria"`
	} `xml:"RequestBody"`
}
