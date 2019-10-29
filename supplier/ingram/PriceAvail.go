package ingram

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"myapp/supplier/model"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

const PA_URL = "https://newport.ingrammicro.com/mustang"

func (c *Client) GetPriceAvail(skus []string) *model.PriceAvailResult {
	var in PNARequest

	in.Version = "2.0"
	in.ShowDetail = "2"
	in.Header.SenderID = "ME"
	in.Header.ReceiverID = "YOU"
	in.Header.CountryCode = "FT"
	in.Header.LoginID = c.Username
	in.Header.Password = c.Password
	in.Header.TransactionID = "1"

	for _, sku := range skus {
		sku = strings.TrimPrefix(sku, "ING-")
		item := RequestItem{SKU: sku}
		in.Items = append(in.Items, item)
	}

	// new http request
	buf, _ := xml.MarshalIndent(in, "", "  ")
	body := bytes.NewBuffer(buf)

	// new http request
	req, err := http.NewRequest("POST", PA_URL, body)
	if err != nil {
		c.LogError(err)
	}

	// http header
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	// http client
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// do request
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		c.LogError(err)
	}

	// parse response
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.LogError(err)
	}

	var x PNAResponse

	reader := bytes.NewReader(response)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&x)
	if err != nil {
		c.LogError(err)
	}

	// logger
	c.LogData(skus[0], PA_URL, string(buf), string(response))

	return ToPriceAvailResult(&x)
}

func ToPriceAvailResult(x *PNAResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
