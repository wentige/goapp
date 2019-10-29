package asi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"myapp/supplier/model"
	"net/http"
	"strings"
	"time"
)

const PA_URL = "https://www.asipartner.com/partneraccess/xml/price.asp"

func (c *Client) GetPriceAvail(skus []string) *model.PriceAvailResult {
	cid := c.Username
	cert := c.Password
	sku := strings.TrimPrefix(skus[0], "AS-") // one SKU per request

	// build URL
	param := fmt.Sprintf("?CID=%s&CERT=%s&SKU=%s", cid, cert, sku)
	url := PA_URL + param

	// new http request
	req, err := http.NewRequest("GET", url, nil)
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

	var x ASIInventory
	err = xml.Unmarshal(response, &x)
	if err != nil {
		c.LogError(err)
	}

	// logger
	c.LogData(sku, PA_URL, param, string(response))

	return ToPriceAvailResult(&x)
}

func ToPriceAvailResult(x *ASIInventory) *model.PriceAvailResult {
	r := &model.PriceAvailResult{}
	return r
}
