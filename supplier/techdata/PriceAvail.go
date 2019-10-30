package techdata

import (
	"bytes"
	"encoding/xml"
	"strings"

	"myapp/supplier/model"
)

const PA_URL = "https://tdxml.techdata.com/xmlservlet"

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

	var x XMLPriceAvailResponse

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
	var x XMLPriceAvailSubmit

	x.Header.Username = self.client.Username
	x.Header.Password = self.client.Password
	x.Header.Version = "1.4"

	for _, sku := range skus {
		sku = strings.TrimPrefix(sku, "TD-")
		info := RequestLineInfo{"VP", sku}
		x.Detail.LineInfo = append(x.Detail.LineInfo, info)
	}

	// convert to xml
	body, _ := xml.MarshalIndent(x, "", "  ")
	return body
}

func (self PriceAvail) SendRequest(url string, body []byte) ([]byte, error) {
	return self.client.SendRequest("POST", url, bytes.NewBuffer(body))
}

func (self PriceAvail) ToPriceAvailResult(x *XMLPriceAvailResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
