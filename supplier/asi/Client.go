package asi

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"myapp/supplier/model"
)

type Config struct {
}

type Client struct {
	Username string
	Password string
	//Logger

	LastRequest  string
	LastResponse string
}

func (c *Client) LogError(err error) {
	log.Println(err)
}

func (c *Client) LogData(sku, url, req, res string) {
	log.Println(sku)
	log.Println(url)
	log.Println(req)
	log.Println(res)
}

func (c *Client) GetPriceAvail(skus []string) *model.PriceAvailResult {
	pa := &PriceAvail{c}
	return pa.Query(skus)
}

func (c *Client) Purchase(item map[string]string) *model.PurchaseResult {
	po := &PurchaseOrder{c}
	return po.Perform(item)
}

func (c *Client) SendRequest(method, url string, data []byte) ([]byte, error) {
	// new http request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
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

	c.LastResponse = string(response)

	return response, nil
}
