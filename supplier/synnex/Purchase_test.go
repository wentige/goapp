package synnex

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PO_Request(t *testing.T) {
	assert := assert.New(t)

	var x OrderRequest

	x.Credential.UserID = "tom@betanade.com"
	x.Credential.Password = "betSyn!2O19053O"

	x.Detail.CustomerNumber = "1150897"
	x.Detail.PONumber = "702-6452671-8794631"
	x.Detail.DropShipFlag = "Y"
	x.Detail.Shipment.ShipFromWarehouse = "57"
	x.Detail.Shipment.ShipTo.AddressName1 = "Lisa Wall"
	x.Detail.Shipment.ShipTo.AddressLine1 = "8 Esterbrooke Ave. TH 23"
	x.Detail.Shipment.ShipTo.AddressLine2 = ""
	x.Detail.Shipment.ShipTo.City = "Toronto"
	x.Detail.Shipment.ShipTo.State = "ON"
	x.Detail.Shipment.ShipTo.ZipCode = "M2J 2C2"
	x.Detail.Shipment.ShipTo.Country = "CA"
	x.Detail.Shipment.ShipToContact.ContactName = "Lisa Wall"
	x.Detail.Shipment.ShipToContact.PhoneNumber = "416 8459399"
	x.Detail.Shipment.ShipToContact.EmailAddress = "pp9qvc9r3gllqv2@marketplace.amazon.ca"
	x.Detail.Shipment.ShipMethod.Code = "PUX"
	x.Detail.Items.Item.LineNumber = "!"
	x.Detail.Items.Item.SKU = "6160176"
	x.Detail.Items.Item.UnitPrice = "61.10"
	x.Detail.Items.Item.Quantity = "1"

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z OrderRequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(z.Credential.UserID, x.Credential.UserID)
	assert.Equal(z.Credential.Password, x.Credential.Password)
	assert.Equal(z.Detail.CustomerNumber, x.Detail.CustomerNumber)
	assert.Equal(z.Detail.PONumber, x.Detail.PONumber)
	assert.Equal(z.Detail.DropShipFlag, x.Detail.DropShipFlag)
	assert.Equal(z.Detail.Shipment.ShipFromWarehouse, x.Detail.Shipment.ShipFromWarehouse)
	assert.Equal(z.Detail.Shipment.ShipTo.AddressName1, x.Detail.Shipment.ShipTo.AddressName1)
	assert.Equal(z.Detail.Shipment.ShipTo.AddressLine1, x.Detail.Shipment.ShipTo.AddressLine1)
	assert.Equal(z.Detail.Shipment.ShipTo.AddressLine2, x.Detail.Shipment.ShipTo.AddressLine2)
	assert.Equal(z.Detail.Shipment.ShipTo.City, x.Detail.Shipment.ShipTo.City)
	assert.Equal(z.Detail.Shipment.ShipTo.State, x.Detail.Shipment.ShipTo.State)
	assert.Equal(z.Detail.Shipment.ShipTo.ZipCode, x.Detail.Shipment.ShipTo.ZipCode)
	assert.Equal(z.Detail.Shipment.ShipTo.Country, x.Detail.Shipment.ShipTo.Country)
	assert.Equal(z.Detail.Shipment.ShipToContact.ContactName, x.Detail.Shipment.ShipToContact.ContactName)
	assert.Equal(z.Detail.Shipment.ShipToContact.PhoneNumber, x.Detail.Shipment.ShipToContact.PhoneNumber)
	assert.Equal(z.Detail.Shipment.ShipToContact.EmailAddress, x.Detail.Shipment.ShipToContact.EmailAddress)
	assert.Equal(z.Detail.Shipment.ShipMethod.Code, x.Detail.Shipment.ShipMethod.Code)
	assert.Equal(z.Detail.Items.Item.LineNumber, x.Detail.Items.Item.LineNumber)
	assert.Equal(z.Detail.Items.Item.SKU, x.Detail.Items.Item.SKU)
	assert.Equal(z.Detail.Items.Item.UnitPrice, x.Detail.Items.Item.UnitPrice)
	assert.Equal(z.Detail.Items.Item.Quantity, x.Detail.Items.Item.Quantity)
}

func Test_PO_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-1.xml"
	data := getPOXML(filename)

	var x OrderResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Detail.Code, "accepted")
	assert.Equal(x.Detail.CustomerNumber, "1150897")
	assert.Equal(x.Detail.PONumber, "702-6452671-8794631")
	assert.Equal(x.Detail.Items.Item.SKU, "6160176")
	assert.Equal(x.Detail.Items.Item.Code, "accepted")
	assert.Equal(x.Detail.Items.Item.OrderNumber, "27169270")
}

func Test_PO_Errors(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
}

func Test_PurchaseOrder(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("SYN_USER")
	client.Password = os.Getenv("SYN_PASS")
	//client.Purchase([]string{"209435"})
}

func getPOXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
