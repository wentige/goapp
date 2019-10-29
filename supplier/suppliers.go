package supplier

import (
	"myapp/supplier/asi"
	"myapp/supplier/dh"
	"myapp/supplier/ingram"
	"myapp/supplier/model"
	"myapp/supplier/synnex"
	"myapp/supplier/techdata"
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
		client = &ingram.Client{}
	case "SYN":
		client = &synnex.Client{}
	case "TD":
		client = &techdata.Client{}
	}

	return client
}
