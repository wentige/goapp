package techdata

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pr = fmt.Println

func Test_Request(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	var x XMLPriceAvailSubmit

	x.Header.Version = "567861"
	x.Header.Username = "btexml9014"
	x.Header.Password = "1.4"
	x.Detail.LineInfo.RefIDQual = "VP"
	x.Detail.LineInfo.RefID = "1568CA"

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z XMLPriceAvailSubmit

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(x.Header.Version, z.Header.Version)
	assert.Equal(x.Header.Username, z.Header.Username)
	assert.Equal(x.Header.Password, z.Header.Password)

	assert.Equal(x.Detail.LineInfo.RefIDQual, z.Detail.LineInfo.RefIDQual)
	assert.Equal(x.Detail.LineInfo.RefID, z.Detail.LineInfo.RefID)
}

func Test_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getXML(filename)

	var x XMLPriceAvailResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Header.TransSetIDCode, "846REC")
	assert.Equal(x.Header.ResponseVersion, "1.4")

	assert.Equal(x.Detail.LineInfo.RefIDQual1, "VP")
	assert.Equal(x.Detail.LineInfo.RefID1, "1568CA")
	assert.Equal(x.Detail.LineInfo.RefIDQual2, "MG")
	assert.Equal(x.Detail.LineInfo.RefID2, "3610/DN")
	assert.Equal(x.Detail.LineInfo.RefIDQual4, "UP")
	assert.Equal(x.Detail.LineInfo.RefID4, "095205972900")
	assert.Equal(x.Detail.LineInfo.ItemStatus, "ACTIVE")
	assert.Equal(x.Detail.LineInfo.ProductDesc, "PHASER 3610 LASERPR 47PPM USB ETH")
	assert.Equal(x.Detail.LineInfo.ProductWeight, "35.200")
	assert.Equal(x.Detail.LineInfo.EndUserInfo, "N")
	assert.Equal(x.Detail.LineInfo.LicenseInfo, "N")

	assert.Equal(len(x.Detail.LineInfo.Warehouses), 2)

	assert.Equal(x.Detail.LineInfo.Warehouses[0].Name, "MISSISSAUGA, ON")
	assert.Equal(x.Detail.LineInfo.Warehouses[0].Code, "A1")
	assert.Equal(x.Detail.LineInfo.Warehouses[0].Qty, "2")

	assert.Equal(x.Detail.LineInfo.Warehouses[1].Name, "RICHMOND, BC")
	assert.Equal(x.Detail.LineInfo.Warehouses[1].Code, "A2")
	assert.Equal(x.Detail.LineInfo.Warehouses[1].Qty, "0")
}

func Test_Errors(t *testing.T) {
}

func Test_Client(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("TD_USER")
	client.Password = os.Getenv("TD_PASS")
	client.GetPriceAvail([]string{"0705XT"})
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
