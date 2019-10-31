package dh

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PO_Request(t *testing.T) {
	assert := assert.New(t)

	var x XmlOrderEntry

	x.Request = "orderEntry"
	x.Login.UserId = "811712XML"
	x.Login.Password = "MTR@xml9013"

	x.Header.OnlyBranch = "Toronto"
	x.Header.Branches = ""
	x.Header.PartshipAllow = "N"
	x.Header.BackorderAllow = "N"
	x.Header.DropshipPW = ""
	x.Header.ShipToName = "Town of Tillsonburg"
	x.Header.ShipToAttn = "(Tel:5196883009)"
	x.Header.ShipToAddress = "10 Lisgar Avenue"
	x.Header.ShipToAddress2 = ""
	x.Header.ShipToCity = "Tillsonburg"
	x.Header.ShipToProvince = "ON"
	x.Header.ShipToPostalCode = "N4G 5A5"
	x.Header.ShipCarrier = "Purolator"
	x.Header.ShipService = "Ground"
	x.Header.PoNum = "702-6223413-8936256"
	x.Header.Remarks = "Drop ship"

	x.Items.Item.Partnum = "8630501CA"
	x.Items.Item.Qty = "1"
	x.Items.Item.Branch = "Toronto"

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z XmlOrderEntry

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(z.Request, x.Request)
	assert.Equal(z.Login.UserId, x.Login.UserId)
	assert.Equal(z.Login.Password, x.Login.Password)

	assert.Equal(z.Header.OnlyBranch, x.Header.OnlyBranch)
	assert.Equal(z.Header.Branches, x.Header.Branches)
	assert.Equal(z.Header.PartshipAllow, x.Header.PartshipAllow)
	assert.Equal(z.Header.BackorderAllow, x.Header.BackorderAllow)
	assert.Equal(z.Header.DropshipPW, x.Header.DropshipPW)
	assert.Equal(z.Header.ShipToName, x.Header.ShipToName)
	assert.Equal(z.Header.ShipToAttn, x.Header.ShipToAttn)
	assert.Equal(z.Header.ShipToAddress, x.Header.ShipToAddress)
	assert.Equal(z.Header.ShipToAddress2, x.Header.ShipToAddress2)
	assert.Equal(z.Header.ShipToCity, x.Header.ShipToCity)
	assert.Equal(z.Header.ShipToProvince, x.Header.ShipToProvince)
	assert.Equal(z.Header.ShipToPostalCode, x.Header.ShipToPostalCode)
	assert.Equal(z.Header.ShipCarrier, x.Header.ShipCarrier)
	assert.Equal(z.Header.ShipService, x.Header.ShipService)
	assert.Equal(z.Header.PoNum, x.Header.PoNum)
	assert.Equal(z.Header.Remarks, x.Header.Remarks)

	assert.Equal(z.Items.Item.Partnum, x.Items.Item.Partnum)
	assert.Equal(z.Items.Item.Qty, x.Items.Item.Qty)
	assert.Equal(z.Items.Item.Branch, x.Items.Item.Branch)
}

func Test_PO_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-1.xml"
	data := getPOXML(filename)

	var x XmlOrderResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Status, "success")
	assert.Equal(x.OrderNum, "4269785")
	assert.Equal(x.Message, "")
}

func Test_PO_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-Error-1.xml"
	data := getPOXML(filename)

	var x XmlOrderResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Status, "failure")
	assert.Equal(x.OrderNum, "")
	assert.Equal(x.Message, "ERROR: request USERID is missing or empty")
}

func Test_PurchaseOrder(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("DH_USER")
	client.Password = os.Getenv("DH_PASS")
	//client.Purchase([]string{"209435"})
}

func getPOXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
