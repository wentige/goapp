package supplier

import (
	"myapp/supplier/asi"
	"myapp/supplier/dh"
	"myapp/supplier/model"
	"strings"
)

//type FileLogger = filelogger.FileLogger

type SupplierClient interface {
	SetLogger(logger map[string]interface{}) // filelogger.FileLogger
	SetConfig(config map[string]interface{})
	GetPriceAvail(skus []string) *model.PriceAvailResult
	PurchaseOrder(items []model.PurchaseItem) *model.PurchaseResult
	//Dropship(...)
	//GetOrderStatus(...)
}

func GetClient(name string) *SupplierClient {
	client := nil

	name = strings.ToUpper(name)
	switch name {
	case "AS":
		client = &asi.Client{}
	case "DH":
		client = &dh.Client{}
	case "ING":
		client = &ing.Client{}
	case "SYN":
		client = &syn.Client{}
	case "TD":
		client = &td.Client{}
	}

	return client
}
