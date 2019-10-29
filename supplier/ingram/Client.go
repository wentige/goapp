package ingram

import (
	"log"
	"myapp/supplier/model"
)

type Config struct {
}

type Client struct {
	Username string
	Password string
	//Logger
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
