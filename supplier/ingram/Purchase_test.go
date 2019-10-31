package ingram

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html/charset"
)

func Test_PO_Request(t *testing.T) {
	assert := assert.New(t)

	var x OrderRequest

	x.Version = "2.0"
	x.ShowDetail = "2"

	x.Header.SenderID = ""
	x.Header.ReceiverID = ""
	x.Header.CountryCode = "FT"
	x.Header.LoginID = "TreV8Efbfe"
	x.Header.Password = "FVnxb28501"
	x.Header.TransactionID = ""

	x.OrderHeader.BillToSuffix = ""
	x.OrderHeader.AddressingInfo.CustomerPO = "70270572879976235"
	x.OrderHeader.AddressingInfo.ShipToAttn = ""
	x.OrderHeader.AddressingInfo.EndUserPO = "70270572879976235"
	x.OrderHeader.AddressingInfo.ShipTo.Address.Address1 = "Troy Martin"
	x.OrderHeader.AddressingInfo.ShipTo.Address.Address2 = "106 Keystone Lane"
	x.OrderHeader.AddressingInfo.ShipTo.Address.Address3 = "(+1 07807392000)"
	x.OrderHeader.AddressingInfo.ShipTo.Address.City = "Leduc"
	x.OrderHeader.AddressingInfo.ShipTo.Address.Province = "AB"
	x.OrderHeader.AddressingInfo.ShipTo.Address.PostalCode = "T9E 0J4"

	x.OrderHeader.ProcessingOptions.CarrierCode = "PI"
	x.OrderHeader.ProcessingOptions.AutoRelease = "H"

	x.OrderHeader.ProcessingOptions.ShipmentOptions.BackOrderFlag = "Y"
	x.OrderHeader.ProcessingOptions.ShipmentOptions.SplitShipmentFlag = "N"
	x.OrderHeader.ProcessingOptions.ShipmentOptions.SplitLine = "N"
	x.OrderHeader.ProcessingOptions.ShipmentOptions.ShipFromBranches = "10"
	x.OrderHeader.ProcessingOptions.ShipmentOptions.DeliveryDate = ""

	x.OrderHeader.DynamicMessage.MessageLines = ""

	x.OrderLine.ProductLine.SKU = "5161UD"
	x.OrderLine.ProductLine.Quantity = "1"
	x.OrderLine.ProductLine.CustomerLineNumber = ""

	x.OrderLine.CommentLine.CommentText = ""

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z OrderRequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(z.Version, x.Version)
	assert.Equal(z.ShowDetail, x.ShowDetail)

	assert.Equal(z.Header.SenderID, x.Header.SenderID)
	assert.Equal(z.Header.ReceiverID, x.Header.ReceiverID)
	assert.Equal(z.Header.CountryCode, x.Header.CountryCode)
	assert.Equal(z.Header.LoginID, x.Header.LoginID)
	assert.Equal(z.Header.Password, x.Header.Password)
	assert.Equal(z.Header.TransactionID, x.Header.TransactionID)

	assert.Equal(z.OrderHeader.BillToSuffix, x.OrderHeader.BillToSuffix)
	assert.Equal(z.OrderHeader.AddressingInfo.CustomerPO, x.OrderHeader.AddressingInfo.CustomerPO)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipToAttn, x.OrderHeader.AddressingInfo.ShipToAttn)
	assert.Equal(z.OrderHeader.AddressingInfo.EndUserPO, x.OrderHeader.AddressingInfo.EndUserPO)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipTo.Address.Address1, x.OrderHeader.AddressingInfo.ShipTo.Address.Address1)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipTo.Address.Address2, x.OrderHeader.AddressingInfo.ShipTo.Address.Address2)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipTo.Address.Address3, x.OrderHeader.AddressingInfo.ShipTo.Address.Address3)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipTo.Address.City, x.OrderHeader.AddressingInfo.ShipTo.Address.City)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipTo.Address.Province, x.OrderHeader.AddressingInfo.ShipTo.Address.Province)
	assert.Equal(z.OrderHeader.AddressingInfo.ShipTo.Address.PostalCode, x.OrderHeader.AddressingInfo.ShipTo.Address.PostalCode)

	assert.Equal(z.OrderHeader.ProcessingOptions.CarrierCode, x.OrderHeader.ProcessingOptions.CarrierCode)
	assert.Equal(z.OrderHeader.ProcessingOptions.AutoRelease, x.OrderHeader.ProcessingOptions.AutoRelease)

	assert.Equal(z.OrderHeader.ProcessingOptions.ShipmentOptions.BackOrderFlag, x.OrderHeader.ProcessingOptions.ShipmentOptions.BackOrderFlag)
	assert.Equal(z.OrderHeader.ProcessingOptions.ShipmentOptions.SplitShipmentFlag, x.OrderHeader.ProcessingOptions.ShipmentOptions.SplitShipmentFlag)
	assert.Equal(z.OrderHeader.ProcessingOptions.ShipmentOptions.SplitLine, x.OrderHeader.ProcessingOptions.ShipmentOptions.SplitLine)
	assert.Equal(z.OrderHeader.ProcessingOptions.ShipmentOptions.ShipFromBranches, x.OrderHeader.ProcessingOptions.ShipmentOptions.ShipFromBranches)
	assert.Equal(z.OrderHeader.ProcessingOptions.ShipmentOptions.DeliveryDate, x.OrderHeader.ProcessingOptions.ShipmentOptions.DeliveryDate)

	assert.Equal(z.OrderHeader.DynamicMessage.MessageLines, x.OrderHeader.DynamicMessage.MessageLines)

	assert.Equal(z.OrderLine.ProductLine.SKU, x.OrderLine.ProductLine.SKU)
	assert.Equal(z.OrderLine.ProductLine.Quantity, x.OrderLine.ProductLine.Quantity)
	assert.Equal(z.OrderLine.ProductLine.CustomerLineNumber, x.OrderLine.ProductLine.CustomerLineNumber)

	assert.Equal(z.OrderLine.CommentLine.CommentText, x.OrderLine.CommentLine.CommentText)
}

func Test_PO_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-1.xml"
	data := getPOXML(filename)

	var x OrderResponse

	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&x)
	assert.Nil(err)

	// There are so many fields in response, we only need a few
	assert.Equal(x.Header.ErrorStatus.ErrorNumber, "")
	assert.Equal(x.OrderInfo.OrderNumbers.BranchOrderNumber, "40G1NPF")
}

func Test_PO_Errors(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
}

func Test_PurchaseOrder(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("ING_USER")
	client.Password = os.Getenv("ING_PASS")
	//client.Purchase([]string{"209435"})
}

func getPOXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
