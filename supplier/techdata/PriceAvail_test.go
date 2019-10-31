package techdata

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PA_Request(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	var x XMLPriceAvailSubmit

	x.Header.Version = "567861"
	x.Header.Username = "btexml9014"
	x.Header.Password = "1.4"
	x.Detail.LineInfo = append(x.Detail.LineInfo, RequestLineInfo{"VP", "1568CA"})

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z XMLPriceAvailSubmit

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(x.Header.Version, z.Header.Version)
	assert.Equal(x.Header.Username, z.Header.Username)
	assert.Equal(x.Header.Password, z.Header.Password)

	assert.Equal(x.Detail.LineInfo[0].RefIDQual, z.Detail.LineInfo[0].RefIDQual)
	assert.Equal(x.Detail.LineInfo[0].RefID, z.Detail.LineInfo[0].RefID)
}

func Test_PA_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getPAXML(filename)

	var x XMLPriceAvailResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Header.TransSetIDCode, "846REC")
	assert.Equal(x.Header.ResponseVersion, "1.4")

	assert.Equal(x.Detail.LineInfo[0].RefIDQual1, "VP")
	assert.Equal(x.Detail.LineInfo[0].RefID1, "1568CA")
	assert.Equal(x.Detail.LineInfo[0].RefIDQual2, "MG")
	assert.Equal(x.Detail.LineInfo[0].RefID2, "3610/DN")
	assert.Equal(x.Detail.LineInfo[0].RefIDQual4, "UP")
	assert.Equal(x.Detail.LineInfo[0].RefID4, "095205972900")
	assert.Equal(x.Detail.LineInfo[0].ItemStatus, "ACTIVE")
	assert.Equal(x.Detail.LineInfo[0].ProductDesc, "PHASER 3610 LASERPR 47PPM USB ETH")
	assert.Equal(x.Detail.LineInfo[0].ProductWeight, "35.200")
	assert.Equal(x.Detail.LineInfo[0].EndUserInfo, "N")
	assert.Equal(x.Detail.LineInfo[0].LicenseInfo, "N")

	assert.Equal(len(x.Detail.LineInfo[0].Warehouses), 2)

	assert.Equal(x.Detail.LineInfo[0].Warehouses[0].Name, "MISSISSAUGA, ON")
	assert.Equal(x.Detail.LineInfo[0].Warehouses[0].Code, "A1")
	assert.Equal(x.Detail.LineInfo[0].Warehouses[0].Qty, "2")

	assert.Equal(x.Detail.LineInfo[0].Warehouses[1].Name, "RICHMOND, BC")
	assert.Equal(x.Detail.LineInfo[0].Warehouses[1].Code, "A2")
	assert.Equal(x.Detail.LineInfo[0].Warehouses[1].Qty, "0")
}

func Test_PA_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-Error-1.xml"
	data := getPOXML(filename)

	var x XMLPriceAvailResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Detail.LineInfo[0].ErrorInfo.ErrorDesc, "The required UserName tag is either blank or missing")
	assert.Equal(x.Detail.LineInfo[0].ErrorInfo.RefIDQual3, "1Q")
	assert.Equal(x.Detail.LineInfo[0].ErrorInfo.RefID3, "0")
}

func Test_GetPriceAvail(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("TD_USER")
	client.Password = os.Getenv("TD_PASS")
	//client.GetPriceAvail([]string{"TD-0705XT", "TD-5665BX"})
}

func getPAXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
