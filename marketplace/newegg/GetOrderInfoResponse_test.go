package newegg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

var pr = fmt.Println
var ppr = pretty.Print

func Test_GetOrderInfoResponse(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/GetOrderInfoResponse-1.xml"
	bytes := getXML(filename)

	var resp GetOrderInfoResponse
	xml.Unmarshal(bytes, &resp)
	//ppr(resp)

	// check pageInfo
	pageInfo := resp.ResponseBody.PageInfo

	assert.Equal(pageInfo.TotalCount, "1")
	assert.Equal(pageInfo.TotalPageCount, "1")
	assert.Equal(pageInfo.PageIndex, "1")
	assert.Equal(pageInfo.PageSize, "100")

	// check orderInfo
	orderInfo := resp.ResponseBody.OrderInfoList.OrderInfo

	assert.Equal(orderInfo.SellerID, "AD6H")
	assert.Equal(orderInfo.OrderNumber, "274286889")

	// check itemInfo
	itemInfo := orderInfo.ItemInfoList.ItemInfo

	assert.Equal(len(itemInfo), 2)
	assert.Equal(itemInfo[0].UPCCode, "0882658345371")
	assert.Equal(itemInfo[0].SellerPartNumber, "ING-48298N")
	assert.Equal(itemInfo[1].UPCCode, "882658345371")
	assert.Equal(itemInfo[1].SellerPartNumber, "DH-SPA303G1CN")

	// check packageInfo
	packageInfo := orderInfo.PackageInfoList.PackageInfo

	assert.Equal(len(packageInfo), 1)
	assert.Equal(packageInfo[0].PackageType, "Shipped")
	assert.Equal(packageInfo[0].ShipCarrier, "UPS")
	assert.Equal(packageInfo[0].TrackingNumber, "1Z37Y0596767691812")
}

func Test_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Errors.xml"
	bytes := getXML(filename)

	var resp Errors
	xml.Unmarshal(bytes, &resp)
	//ppr(resp)

	assert.Equal(resp.Error[0].Code, "InvalidToken")
	assert.Equal(resp.Error[0].Message, "The provided secret key is null.")
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
