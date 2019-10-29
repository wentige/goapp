package asi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"myapp/supplier/model"
	"net/http"
	"strings"
	"time"
)

const PO_URL = ""

type PurchaseOrder struct {
	client *Client
}

func (self PurchaseOrder) Perform(items []model.PurchaseItem) *model.PurchaseResult {
	request := self.BuildRequest(items)

	response, err := self.SendRequest(PO_URL, request)
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
	self.client.LogData(skus[0], PO_URL, string(request), string(response))

	return self.ToPurchaseResult(&x)
}

func (self PurchaseOrder) BuildRequest(skus []string) string {
	cid := self.client.Username
	cert := self.client.Password
	sku := strings.TrimPrefix(skus[0], "AS-") // one SKU per request

	return fmt.Sprintf("?CID=%s&CERT=%s&SKU=%s", cid, cert, sku)
}

func (self PurchaseOrder) SendRequest(url string, body []byte) ([]byte, error) {
	// new http request
	req, err := http.NewRequest("POST", url, nil)
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
	r := &model.PurchaseOrderResult{}
	return r
}
