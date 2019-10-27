package synnex

import (
	"encoding/xml"
)

// Request

type SkuInfo struct {
	SKU        string `xml:"synnexSKU"`
	LineNumber int    `xml:"lineNumber"`
}

type PriceRequest struct {
	XMLName    xml.Name  `xml:"priceRequest"`
	CustomerNo string    `xml:"customerNo"`
	Username   string    `xml:"userName"`
	Password   string    `xml:"password"`
	SkuList    []SkuInfo `xml:"skuList"`
}

// Response

type WarehouseInfo struct {
	Number  string `xml:"number"`
	Zipcode string `xml:"zipcode"`
	City    string `xml:"city"`
	Addr    string `xml:"addr"`
}

type Warehouse struct {
	Info WarehouseInfo `xml:"warehouseInfo"`
	Qty  string        `xml:"qty"`
}

type ItemInfo struct {
	SKU         string      `xml:"synnexSKU"`
	MfgPN       string      `xml:"mfgPN"`
	MfgCode     string      `xml:"mfgCode"`
	Status      string      `xml:"status"`
	Description string      `xml:"description"`
	StatusCode  string      `xml:"GlobalProductStatusCode"`
	Price       string      `xml:"price"`
	TotalQty    string      `xml:"totalQuantity"`
	Warehouses  []Warehouse `xml:"AvailabilityByWarehouse"`
	LineNumber  string      `xml:"lineNumber"`
}

type PriceResponse struct {
	XMLName    xml.Name   `xml:"priceResponse"`
	CustomerNo string     `xml:"customerNo"`
	Username   string     `xml:"userName"`
	Items      []ItemInfo `xml:"PriceAvailabilityList"`
}
