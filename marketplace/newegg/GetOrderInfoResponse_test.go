package newegg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
)

var pr = fmt.Println
var ppr = pretty.Print

func Test_GetOrderInfoResponse(t *testing.T) {
	filename := "fixtures/GetOrderInfoResponse-1.xml"
	bytes := getXML(filename)

	var resp GetOrderInfoResponse
	xml.Unmarshal(bytes, &resp)
	//ppr(resp)

	// check pageInfo
	pageInfo := resp.ResponseBody.PageInfo

	if pageInfo.TotalCount != "1" {
		t.Error("TotalCount not expected")
	}
	if pageInfo.TotalPageCount != "1" {
		t.Error("TotalPageCount not expected")
	}
	if pageInfo.PageIndex != "1" {
		t.Error("PageIndex not expected")
	}
	if pageInfo.PageSize != "100" {
		t.Error("PageSize not expected")
	}

	// check orderInfo
	orderInfo := resp.ResponseBody.OrderInfoList.OrderInfo

	if orderInfo.SellerID != "AD6H" {
		t.Error("SellerID not expected")
	}
	if orderInfo.OrderNumber != "274286889" {
		t.Error("OrderNumber not expected")
	}

	// check itemInfo
	itemInfo := orderInfo.ItemInfoList.ItemInfo

	if len(itemInfo) != 2 {
		t.Error("ItemInfo[] not expected")
	}
	if itemInfo[0].UPCCode != "0882658345371" {
		t.Error("UPCCode-0 not expected")
	}
	if itemInfo[0].SellerPartNumber != "ING-48298N" {
		t.Error("SellerPartNumber-0 not expected")
	}
	if itemInfo[1].UPCCode != "882658345371" {
		t.Error("UPCCode-1 not expected")
	}
	if itemInfo[1].SellerPartNumber != "DH-SPA303G1CN" {
		t.Error("SellerPartNumber-0 not expected")
	}

	// check packageInfo
	packageInfo := orderInfo.PackageInfoList.PackageInfo

	if len(packageInfo) != 1 {
		t.Error("PackageInfo[] not expected")
	}
	if packageInfo[0].PackageType != "Shipped" {
		t.Error("PackageType not expected")
	}
	if packageInfo[0].ShipCarrier != "UPS" {
		t.Error("ShipCarrier not expected")
	}
	if packageInfo[0].TrackingNumber != "1Z37Y0596767691812" {
		t.Error("TrackingNumber not expected")
	}
}

func Test_Errors(t *testing.T) {
	filename := "fixtures/Errors.xml"
	bytes := getXML(filename)

	var resp Errors
	xml.Unmarshal(bytes, &resp)
	//ppr(resp)

	if resp.Error[0].Code != "InvalidToken" {
		t.Error("Error Code not expected")
	}
	if resp.Error[0].Message != "The provided secret key is null." {
		t.Error("Error Message not expected")
	}
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
