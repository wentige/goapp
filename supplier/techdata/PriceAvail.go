package techdata

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"myapp/supplier/model"
	"net/http"
	"strings"
	"time"
)

const PA_URL = "https://tdxml.techdata.com/xmlservlet"

func (c *Client) GetPriceAvail(skus []string) *model.PriceAvailResult {
	var in XMLPriceAvailSubmit

	in.Header.Username = c.Username
	in.Header.Password = c.Password
	in.Header.Version = "1.4"

	for _, sku := range skus {
		sku = strings.TrimPrefix(sku, "TD-")
		info := RequestLineInfo{"VP", sku}
		in.Detail.LineInfo = append(in.Detail.LineInfo, info)
	}

	// convert to xml
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

	var x XMLPriceAvailResponse
	err = xml.Unmarshal(response, &x)
	if err != nil {
		c.LogError(err)
	}

	// logger
	c.LogData(skus[0], PA_URL, string(buf), string(response))

	return ToPriceAvailResult(&x)
}

func ToPriceAvailResult(x *XMLPriceAvailResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
