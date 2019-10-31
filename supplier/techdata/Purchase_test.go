package techdata

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PO_Request(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	var x XMLOrderSubmit

	x.Header.Username = "567861"
	x.Header.Password = "bettdxml9014"
	x.Header.ResponseVersion = "1.6"
	x.Header.OrderTypeCode = "BS"
	x.Header.PoNum = "701-0093022-1599463"
	x.Header.RefNum.RefIDQual = "EU"
	x.Header.RefNum.RefID = "701-0093022-1599463"
	x.Header.SalesRequirementCode = ""
	x.Header.Name = "Groupe Major"
	x.Header.AddrInfo.Addr = []string{"1255 boul Lebourgneuf", "Suite 280", "418.730.2797"}
	x.Header.CityName = "Quebec"
	x.Header.StateProvinceCode = "QC"
	x.Header.PostalCode = "G2K 0M6"
	x.Header.ContactName = "Roy Zhang"
	x.Header.ContactPhone = "418.730.2797"

	x.Header.EndUserInfo.GovAgency = ""
	x.Header.EndUserInfo.GovCabinetLevel = ""
	x.Header.EndUserInfo.ContractNum = ""
	x.Header.EndUserInfo.ContractType = ""
	x.Header.EndUserInfo.OrderPriority = ""
	x.Header.EndUserInfo.MarketType = ""
	x.Header.EndUserInfo.ContactName = "Groupe Major"
	x.Header.EndUserInfo.PhoneNum = "418.730.2797"
	x.Header.EndUserInfo.FaxNum = ""
	x.Header.EndUserInfo.Name = "Groupe Major"
	x.Header.EndUserInfo.Addr1 = "1255 boul Lebourgneuf Suite 280"
	x.Header.EndUserInfo.Addr2 = ""
	x.Header.EndUserInfo.Addr3 = ""
	x.Header.EndUserInfo.CityName = "Quebec"
	x.Header.EndUserInfo.StateProvinceCode = "QC"
	x.Header.EndUserInfo.PostalCode = "G2K 0M6"
	x.Header.EndUserInfo.CountryCode = "CA"
	x.Header.EndUserInfo.SicCode = ""
	x.Header.EndUserInfo.OrderPromoType = ""
	x.Header.EndUserInfo.EndUserLicenseNum = ""
	x.Header.EndUserInfo.EndUserPODate = ""
	x.Header.EndUserInfo.EndUserRef1 = ""
	x.Header.EndUserInfo.EndUserRef2 = ""
	x.Header.EndUserInfo.EndUserRef3 = ""
	x.Header.EndUserInfo.InstallName = ""
	x.Header.EndUserInfo.DropShipType = ""
	x.Header.EndUserInfo.ContactEmailAddr1 = ""
	x.Header.EndUserInfo.ContactEmailAddr2 = ""
	x.Header.MyOrderTracker.ResellerEmail = "jessica@btecanada.com"
	x.Header.MyOrderTracker.ResellerEvents = []string{"OC", "OS"}

	x.Detail.LineInfo.QtyOrdered = "1"
	x.Detail.LineInfo.ProductIDQual = "VP"
	x.Detail.LineInfo.ProductID = "0331ZE"
	x.Detail.LineInfo.WhseCode = "A1"
	x.Detail.LineInfo.IDCode = "UP"
	x.Detail.LineInfo.OrderMessageLine = ""

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	var z XMLOrderSubmit

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(z.Header.Username, x.Header.Username)
	assert.Equal(z.Header.Password, x.Header.Password)
}

func Test_PO_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-1.xml"
	data := getPOXML(filename)

	var x XMLOrderResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Header.Username, "567861")
	assert.Equal(x.Header.PONum, "701-0093022-1599463")
	assert.Equal(x.Header.RefID, "1315253")
}

func Test_PO_Errors(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
}

func Test_PurchaseOrder(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("TD_USER")
	client.Password = os.Getenv("TD_PASS")
	//client.Purchase([]string{"209435"})
}

func getPOXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
