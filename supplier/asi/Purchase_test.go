package asi

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PO_Request(t *testing.T) {
	assert := assert.New(t)

	var x ASIOrderRequest

	x.Custid = "75692"
	x.Custpo = "W-1572350427-AS"
	x.Cert = "1FN9RYH3GNDH5OM"
	x.Header.Shipto.Name = "Roy Zhang"
	x.Header.Shipto.Address1 = "133 McPherson Street"
	x.Header.Shipto.Address2 = ""
	x.Header.Shipto.Address3 = "647-400-9097"
	x.Header.Shipto.City = "Markham"
	x.Header.Shipto.State = "ON"
	x.Header.Shipto.Zip = "L3R 3L3"
	x.Header.Shipto.Country = "CA"
	x.Header.Shipinfo.Via = "PGD"
	x.Header.Shipinfo.Instructions = "Please HOLD"
	x.Detail.Line.Sku = "228303"
	x.Detail.Line.Qty = "1"
	x.Detail.Line.Price = "538.23"
	x.Detail.Line.Branch = "3116"

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	var z ASIOrderRequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(z.Custid, x.Custid)
	assert.Equal(z.Custpo, x.Custpo)
	assert.Equal(z.Cert, x.Cert)
	assert.Equal(z.Header.Shipto.Name, x.Header.Shipto.Name)
	assert.Equal(z.Header.Shipto.Address1, x.Header.Shipto.Address1)
	assert.Equal(z.Header.Shipto.Address2, x.Header.Shipto.Address2)
	assert.Equal(z.Header.Shipto.Address3, x.Header.Shipto.Address3)
	assert.Equal(z.Header.Shipto.City, x.Header.Shipto.City)
	assert.Equal(z.Header.Shipto.State, x.Header.Shipto.State)
	assert.Equal(z.Header.Shipto.Zip, x.Header.Shipto.Zip)
	assert.Equal(z.Header.Shipto.Country, x.Header.Shipto.Country)
	assert.Equal(z.Header.Shipinfo.Via, x.Header.Shipinfo.Via)
	assert.Equal(z.Header.Shipinfo.Instructions, x.Header.Shipinfo.Instructions)
	assert.Equal(z.Detail.Line.Sku, x.Detail.Line.Sku)
	assert.Equal(z.Detail.Line.Qty, x.Detail.Line.Qty)
	assert.Equal(z.Detail.Line.Price, x.Detail.Line.Price)
	assert.Equal(z.Detail.Line.Branch, x.Detail.Line.Branch)
}

func Test_PO_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-1.xml"
	data := getPOXML(filename)

	var x ASIOrderReply

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)
	assert.Equal(x.Order.Orderid, "12451302")
	assert.Equal(x.Error.Code, "")
	assert.Equal(x.Error.Message, "")
}

func Test_PO_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-Error-1.xml"
	data := getPOXML(filename)

	var x ASIOrderReply

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)
	assert.Equal(x.Order.Orderid, "")
	assert.Equal(x.Error.Code, "101")
	assert.Equal(x.Error.Message, "SKU 100166's price does not match your XML feed price. It should be $74.08")
}

func Test_PurchaseOrder(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("ASI_USER")
	client.Password = os.Getenv("ASI_PASS")
	//client.Purchase()
}

func getPOXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
