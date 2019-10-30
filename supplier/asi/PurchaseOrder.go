package asi

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"myapp/supplier/model"
	"net/http"
	"time"
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
	// new http request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// http header
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	// http client
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// send request
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	// read response
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (self PurchaseOrder) ToPurchaseResult(x *ASIInventory) *model.PurchaseResult {
	r := &model.PurchaseResult{}
	return r
}
