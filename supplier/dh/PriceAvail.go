package dh

import (
	"encoding/xml"
	"strings"

	"myapp/supplier/model"
)

const PA_URL = "https://www.dandh.ca/dhXML/xmlDispatch"

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

	var x XmlResponse

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
	var x XmlRequest

	x.Request = "price-availability"
	x.Login.UserId = self.client.Username
	x.Login.Password = self.client.Password
	//x.Partnums = make([]string, 0)

	for _, sku := range skus {
		x.Partnums = append(x.Partnums, strings.TrimPrefix(sku, "DH-"))
	}

	body, _ := xml.MarshalIndent(x, "", "  ")
	return body
}

func (self PriceAvail) SendRequest(url string, body []byte) ([]byte, error) {
	return self.client.SendRequest("POST", url, body)
}

func (self PriceAvail) ToPriceAvailResult(x *XmlResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
