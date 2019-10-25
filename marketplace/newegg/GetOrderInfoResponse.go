package newegg

import (
	"encoding/xml"
)

type PageInfo struct {
	TotalCount     string `xml:"TotalCount"`
	TotalPageCount string `xml:"TotalPageCount"`
	PageIndex      string `xml:"PageIndex"`
	PageSize       string `xml:"PageSize"`
}

type ItemInfoList struct {
	ItemInfo []ItemInfo `xml:"ItemInfo"`
}

type ItemInfo struct {
	SellerPartNumber     string `xml:"SellerPartNumber"`
	NeweggItemNumber     string `xml:"NeweggItemNumber"`
	MfrPartNumber        string `xml:"MfrPartNumber"`
	UPCCode              string `xml:"UPCCode"`
	Description          string `xml:"Description"`
	OrderedQty           string `xml:"OrderedQty"`
	ShippedQty           string `xml:"ShippedQty"`
	UnitPrice            string `xml:"UnitPrice"`
	ExtendUnitPrice      string `xml:"ExtendUnitPrice"`
	ExtendShippingCharge string `xml:"ExtendShippingCharge"`
	ExtendSalesTax       string `xml:"ExtendSalesTax"`
	ExtendVAT            string `xml:"ExtendVAT"`
	ExtendDuty           string `xml:"ExtendDuty"`
	Status               string `xml:"Status"`
	StatusDescription    string `xml:"StatusDescription"`
}

type PackageInfo struct {
	PackageType    string       `xml:"PackageType"`
	ShipCarrier    string       `xml:"ShipCarrier"`
	ShipService    string       `xml:"ShipService"`
	TrackingNumber string       `xml:"TrackingNumber"`
	ShipDate       string       `xml:"ShipDate"`
	ItemInfoList   ItemInfoList `xml:"ItemInfoList"`
}

type PackageInfoList struct {
	PackageInfo []PackageInfo `xml:"PackageInfo"`
}

type OrderInfo struct {
	SellerOrderNumber      string          `xml:"SellerOrderNumber"`
	SellerID               string          `xml:"SellerID"`
	OrderNumber            string          `xml:"OrderNumber"`
	InvoiceNumber          string          `xml:"InvoiceNumber"`
	OrderDownloaded        string          `xml:"OrderDownloaded"`
	OrderDate              string          `xml:"OrderDate"`
	OrderStatus            string          `xml:"OrderStatus"`
	OrderStatusDescription string          `xml:"OrderStatusDescription"`
	CustomerName           string          `xml:"CustomerName"`
	CustomerPhoneNumber    string          `xml:"CustomerPhoneNumber"`
	CustomerEmailAddress   string          `xml:"CustomerEmailAddress"`
	ShipToAddress1         string          `xml:"ShipToAddress1"`
	ShipToAddress2         string          `xml:"ShipToAddress2"`
	ShipToCityName         string          `xml:"ShipToCityName"`
	ShipToStateCode        string          `xml:"ShipToStateCode"`
	ShipToZipCode          string          `xml:"ShipToZipCode"`
	ShipToCountryCode      string          `xml:"ShipToCountryCode"`
	ShipService            string          `xml:"ShipService"`
	ShipToFirstName        string          `xml:"ShipToFirstName"`
	ShipToLastName         string          `xml:"ShipToLastName"`
	ShipToCompany          string          `xml:"ShipToCompany"`
	OrderItemAmount        string          `xml:"OrderItemAmount"`
	ShippingAmount         string          `xml:"ShippingAmount"`
	DiscountAmount         string          `xml:"DiscountAmount"`
	RefundAmount           string          `xml:"RefundAmount"`
	SalesTax               string          `xml:"SalesTax"`
	VATTotal               string          `xml:"VATTotal"`
	DutyTotal              string          `xml:"DutyTotal"`
	OrderTotalAmount       string          `xml:"OrderTotalAmount"`
	OrderQty               string          `xml:"OrderQty"`
	IsAutoVoid             string          `xml:"IsAutoVoid"`
	SalesChannel           string          `xml:"SalesChannel"`
	FulfillmentOption      string          `xml:"FulfillmentOption"`
	ItemInfoList           ItemInfoList    `xml:"ItemInfoList"`
	PackageInfoList        PackageInfoList `xml:"PackageInfoList"`
}

type OrderInfoList struct {
	OrderInfo OrderInfo `xml:"OrderInfo"`
}

type ResponseBody struct {
	PageInfo      PageInfo      `xml:"PageInfo"`
	OrderInfoList OrderInfoList `xml:"OrderInfoList"`
}

type GetOrderInfoResponse struct {
	XMLName       xml.Name     `xml:"NeweggAPIResponse"`
	IsSuccess     string       `xml:"IsSuccess"`
	OperationType string       `xml:"OperationType"`
	SellerID      string       `xml:"SellerID"`
	ResponseBody  ResponseBody `xml:"ResponseBody"`
	ResponseDate  string       `xml:"ResponseDate"`
	Memo          string       `xml:"Memo"`
}
