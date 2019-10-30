package synnex

import (
	"bytes"
	"encoding/xml"
	"strings"

	"myapp/supplier/model"
)

const PA_URL = "https://ec.synnex.ca/SynnexXML/PriceAvailability"

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

	var x PriceResponse

	err = xml.Unmarshal(response, &x)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	// logger
	self.client.LogData(skus[0], PA_URL, string(request), string(response))

	return self.ToPriceAvailResult(&x)
}

func (self PriceAvail) BuildRequest(skus []string) []byte {
	var x PriceRequest

	// fill the request struct
	x.CustomerNo = self.client.CustomerNo
	x.Username = self.client.Username
	x.Password = self.client.Password

	for i, sku := range skus {
		sku = strings.TrimPrefix(sku, "SYN-")
		item := SkuInfo{sku, i + 1}
		x.SkuList = append(x.SkuList, item)
	}

	// convert to xml
	body, _ := xml.MarshalIndent(x, "", "  ")
	return body
}

func (self PriceAvail) SendRequest(url string, body []byte) ([]byte, error) {
	return self.client.SendRequest("POST", url, bytes.NewBuffer(body))
}

func (self PriceAvail) ToPriceAvailResult(x *PriceResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
