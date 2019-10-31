package asi

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
}

func Test_PA_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getPAXML(filename)

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

func Test_PA_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-Error-1.xml"
	data := getPAXML(filename)

	var x ASIInventory

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Error.Code, "101")
	assert.Equal(x.Error.Message, "No inventory found")
}

func Test_GetPriceAvail(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("ASI_USER")
	client.Password = os.Getenv("ASI_PASS")
	//client.GetPriceAvail([]string{"209435"})
}

func getPAXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
