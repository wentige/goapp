package asi

import (
	"encoding/xml"

	"myapp/supplier/model"
)

const PO_URL = "https://www.asipartner.com/partneraccess/xml/order.asp"

type PurchaseOrder struct {
	client *Client
}

func (self PurchaseOrder) Perform(item map[string]string) *model.PurchaseResult {
	request := self.BuildRequest(item)

	response, err := self.SendRequest(PO_URL, request)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	var x ASIOrderReply

	err = xml.Unmarshal(response, &x)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	// logger
	sku := item["SKU"]
	self.client.LogData(sku, PO_URL, string(request), string(response))

	return self.ToPurchaseResult(&x)
}

func (self PurchaseOrder) BuildRequest(item map[string]string) []byte {
	var x ASIOrderRequest

	body, _ := xml.MarshalIndent(x, "", "  ")
	return body
}

func (self PurchaseOrder) SendRequest(url string, body []byte) ([]byte, error) {
	return self.client.SendRequest("POST", url, body)
}

func (self PurchaseOrder) ToPurchaseResult(x *ASIOrderReply) *model.PurchaseResult {
	r := &model.PurchaseResult{}
	return r
}
