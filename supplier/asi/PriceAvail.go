package asi

import (
	"encoding/xml"
	"fmt"
	"myapp/supplier/model"
	"strings"
)

const PA_URL = "https://www.asipartner.com/partneraccess/xml/price.asp"

type PriceAvail struct {
	client *Client
}

func (self PriceAvail) Query(skus []string) *model.PriceAvailResult {
	params := self.BuildRequest(skus)

	url := PA_URL + params

	response, err := self.SendRequest(url)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	var x ASIInventory

	err = xml.Unmarshal(response, &x)
	if err != nil {
		self.client.LogError(err)
		return nil
	}

	// logger
	self.client.LogData(skus[0], PA_URL, params, string(response))

	return self.ToPriceAvailResult(&x)
}

func (self PriceAvail) BuildRequest(skus []string) string {
	cid := self.client.Username
	cert := self.client.Password
	sku := strings.TrimPrefix(skus[0], "AS-") // one SKU per request

	return fmt.Sprintf("?CID=%s&CERT=%s&SKU=%s", cid, cert, sku)
}

func (self PriceAvail) SendRequest(url string) ([]byte, error) {
	return self.client.SendRequest("GET", url, nil)
}

func (self PriceAvail) ToPriceAvailResult(x *ASIInventory) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
