package asi

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"myapp/supplier/model"
	"net/http"
	"strings"
	"time"
)

const PA_URL = "https://www.asipartner.com/partneraccess/xml/price.asp"

type PriceAvail struct {
	client *Client
}

func (self PriceAvail) Query(skus []string) *model.PriceAvailResult {
	params := self.BuildRequest(skus)

	url := PA_URL + params

	response, err := self.SendRequest(url, nil)
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

	return ToPriceAvailResult(&x)
}

func (self PriceAvail) BuildRequest(skus []string) string {
	cid := self.client.Username
	cert := self.client.Password
	sku := strings.TrimPrefix(skus[0], "AS-") // one SKU per request

	return fmt.Sprintf("?CID=%s&CERT=%s&SKU=%s", cid, cert, sku)
}

func (self PriceAvail) SendRequest(url string, body io.Reader) ([]byte, error) {
	// new http request
	req, err := http.NewRequest("GET", url, nil)
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

func ToPriceAvailResult(x *ASIInventory) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
