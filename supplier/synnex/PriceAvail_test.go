package synnex

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

	var x PriceRequest

	x.CustomerNo = "11508"
	x.Username = "user@name.com"
	x.Password = "pa55w0rd"

	x.SkuList = []SkuInfo{
		{SKU: "5688596", LineNumber: 1},
	}

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z PriceRequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(x.CustomerNo, z.CustomerNo)
	assert.Equal(x.Username, z.Username)
	assert.Equal(x.Password, z.Password)
	assert.Equal(len(x.SkuList), len(z.SkuList))
	assert.Equal(x.SkuList[0].SKU, z.SkuList[0].SKU)
	assert.Equal(x.SkuList[0].LineNumber, z.SkuList[0].LineNumber)
}

func Test_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getXML(filename)

	var x PriceResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.CustomerNo, "1150897")
	assert.Equal(x.Username, "roi@btxcdn.com")

	assert.Equal(len(x.Items), 1)
	assert.Equal(x.Items[0].SKU, "5688596")
	assert.Equal(x.Items[0].Status, "Active")
	assert.Equal(x.Items[0].TotalQty, "882")

	assert.Equal(len(x.Items[0].Warehouses), 2)

	assert.Equal(x.Items[0].Warehouses[0].Qty, "6")
	assert.Equal(x.Items[0].Warehouses[0].Info.Number, "26")

	assert.Equal(x.Items[0].Warehouses[1].Qty, "476")
	assert.Equal(x.Items[0].Warehouses[1].Info.Number, "29")
}

func Test_Errors(t *testing.T) {
}

func Test_Client(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("SYN_USER")
	client.Password = os.Getenv("SYN_PASS")
	client.GetPriceAvail([]string{"6024136"})
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
