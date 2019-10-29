package asi

import "log"

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
