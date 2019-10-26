package ingram

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	x.Info.SKU = "SKU-1"
	x.Info.Qty = "123"
	x.ShowDetail = "2"

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z PNARequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	if x.Version != z.Version {
		t.Error("Version not expected")
	}
	if x.Header.SenderID != z.Header.SenderID {
		t.Error("Header.SenderID not expected")
	}
	if x.Header.LoginID != z.Header.LoginID {
		t.Error("Header.LoginID not expected")
	}
	if x.Info.SKU != z.Info.SKU {
		t.Error("Info.SKU not expected")
	}
	if x.Info.Qty != z.Info.Qty {
		t.Error("Info.Qty not expected")
	}
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
	if x.Version != "2.0" {
		t.Error("Version not expected")
	}
	if x.Header.SenderID != "YOU" {
		t.Error("Header.SenderID not expected")
	}

	// Item
	if x.Item.SKU != "1553CK" {
		t.Error("Item.SKU not expected")
	}
	if x.Item.Qty != "1" {
		t.Error("Item.Qty not expected")
	}
	if x.Item.UPC != "0763649064979" {
		t.Error("Item.UPC not expected")
	}

	// Branch
	if len(x.Item.Branches) != 2 {
		t.Error("Item.Branches[] not expected")
	}

	if x.Item.Branches[0].ID != "10" {
		t.Error("Item.Branches[0].ID not expected")
	}
	if x.Item.Branches[0].Name != "Vancouver" {
		t.Error("Item.Branches[0].Name not expected")
	}
	if x.Item.Branches[0].Availability != "142" {
		t.Error("Item.Branches[0].Avail not expected")
	}

	if x.Item.Branches[1].ID != "40" {
		t.Error("Item.Branches[1].ID not expected")
	}
	if x.Item.Branches[1].Name != "Toronto" {
		t.Error("Item.Branches[1].Name not expected")
	}
	if x.Item.Branches[1].Availability != "399" {
		t.Error("Item.Branches[1].Avail not expected")
	}
}

func Test_Errors(t *testing.T) {
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
