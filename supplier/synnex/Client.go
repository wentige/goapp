package synnex

import (
	"log"
	"myapp/supplier/model"
)

type Config struct {
}

type Client struct {
	CustomerNo string
	Username   string
	Password   string
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

func (c *Client) Purchase(items []model.PurchaseItem) *model.PurchaseResult {
	po := &PurchaseOrder{c}
	return po.Perform(items)
}
