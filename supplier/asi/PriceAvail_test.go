package asi

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
}

func Test_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getXML(filename)

	var x ASIInventory

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Error.Code, "")

	assert.Equal(len(x.Items), 1)
	assert.Equal(x.Items[0].SKU, "209435")
	assert.Equal(x.Items[0].ItemId, "TS-431XEU-2G-US")
	assert.Equal(x.Items[0].UPC, "885022014033")
	assert.Equal(x.Items[0].Status, "ACTIVE")

	assert.Equal(len(x.Items[0].Qty.Branches), 2)

	assert.Equal(x.Items[0].Qty.Branches[0].Name, "Toronto")
	assert.Equal(x.Items[0].Qty.Branches[0].Code, "2216")
	assert.Equal(x.Items[0].Qty.Branches[0].Qty, "2")

	assert.Equal(x.Items[0].Qty.Branches[1].Name, "Vancouver")
	assert.Equal(x.Items[0].Qty.Branches[1].Code, "3116")
	assert.Equal(x.Items[0].Qty.Branches[1].Qty, "1")
}

func Test_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-Error-1.xml"
	data := getXML(filename)

	var x ASIInventory

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Error.Code, "101")
	assert.Equal(x.Error.Message, "No inventory found")
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
