package ingram

import (
	"bytes"
	"encoding/xml"
	"strings"

	"myapp/supplier/model"

	"golang.org/x/net/html/charset"
)

const PA_URL = "https://newport.ingrammicro.com/mustang"

type PriceAvail struct {
	client *Client
}

func (self PriceAvail) Query(skus []string) *model.PriceAvailResult {
	request := self.BuildRequest(skus)

	response, err := self.SendRequest(PA_URL, request)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	var x PNAResponse

	reader := bytes.NewReader(response)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&x)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	// logger
	self.client.LogData(skus[0], PA_URL, string(request), string(response))

	return self.ToPriceAvailResult(&x)
}

func (self PriceAvail) BuildRequest(skus []string) []byte {
	var x PNARequest

	x.Version = "2.0"
	x.ShowDetail = "2"
	x.Header.SenderID = "ME"
	x.Header.ReceiverID = "YOU"
	x.Header.CountryCode = "FT"
	x.Header.LoginID = self.client.Username
	x.Header.Password = self.client.Password
	x.Header.TransactionID = "1"

	for _, sku := range skus {
		sku = strings.TrimPrefix(sku, "ING-")
		item := RequestItem{SKU: sku}
		x.Items = append(x.Items, item)
	}

	body, _ := xml.MarshalIndent(x, "", "  ")
	return body
}

func (self PriceAvail) SendRequest(url string, body []byte) ([]byte, error) {
	return self.client.SendRequest("POST", url, body)
}

func (self PriceAvail) ToPriceAvailResult(x *PNAResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
