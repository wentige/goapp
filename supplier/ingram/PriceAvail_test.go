package ingram

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html/charset"
)

var pr = fmt.Println

func Test_PNARequest(t *testing.T) {
	assert := assert.New(t)

	var x PNARequest

	x.Version = "2.0"
	x.Header.SenderID = "ME"
	x.Header.ReceiverID = "YOU"
	x.Header.CountryCode = "FT"
	x.Header.LoginID = "LOGIN-ID"
	x.Header.Password = "PASSWORD"
	x.Header.TransactionID = "1"
	x.Items = append(x.Items, RequestItem{"SKU-1", "123"})
	x.ShowDetail = "2"

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z PNARequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(x.Version, z.Version)
	assert.Equal(x.Header.SenderID, z.Header.SenderID)
	assert.Equal(x.Header.LoginID, z.Header.LoginID)
	assert.Equal(x.Items[0].SKU, z.Items[0].SKU)
	assert.Equal(x.Items[0].Qty, z.Items[0].Qty)
}

func Test_PNAResponse(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getXML(filename)

	var x PNAResponse

	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&x)
	assert.Nil(err)

	// Header
	assert.Equal(x.Version, "2.0")
	assert.Equal(x.Header.SenderID, "YOU")

	// Item
	assert.Equal(x.Items[0].SKU, "1553CK")
	assert.Equal(x.Items[0].Qty, "1")
	assert.Equal(x.Items[0].UPC, "0763649064979")

	// Branch
	assert.Equal(len(x.Items[0].Branches), 2)
	assert.Equal(x.Items[0].Branches[0].ID, "10")
	assert.Equal(x.Items[0].Branches[0].Name, "Vancouver")
	assert.Equal(x.Items[0].Branches[0].Availability, "142")

	assert.Equal(x.Items[0].Branches[1].ID, "40")
	assert.Equal(x.Items[0].Branches[1].Name, "Toronto")
	assert.Equal(x.Items[0].Branches[1].Availability, "399")
}

func Test_Errors(t *testing.T) {
}

func Test_Client(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("ING_USER")
	client.Password = os.Getenv("ING_PASS")
	client.GetPriceAvail([]string{"ING-8059YD", "ING-9932DS"})
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
