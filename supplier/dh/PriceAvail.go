package dh

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"myapp/supplier/model"
	"net/http"
	"strings"
	"time"
)

const PA_URL = "https://www.dandh.ca/dhXML/xmlDispatch"

func (c *Client) GetPriceAvail(skus []string) *model.PriceAvailResult {
	var in XmlRequest

	in.Request = "price-availability"
	in.Login.UserId = c.Username
	in.Login.Password = c.Password
	//in.Partnums = make([]string, 0)
	for _, sku := range skus {
		in.Partnums = append(in.Partnums, strings.TrimPrefix(sku, "DH-"))
	}

	buf, _ := xml.MarshalIndent(in, "", "  ")
	body := bytes.NewBuffer(buf)

	// new http request
	req, err := http.NewRequest("POST", PA_URL, body)
	if err != nil {
		log.Println(err)
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

	var x XmlResponse
	err = xml.Unmarshal(response, &x)
	if err != nil {
		c.LogError(err)
	}

	// logger
	c.LogData(skus[0], PA_URL, string(buf), string(response))

	return ToPriceAvailResult(&x)
}

func ToPriceAvailResult(x *XmlResponse) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
