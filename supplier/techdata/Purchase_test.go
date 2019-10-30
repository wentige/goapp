package techdata

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pr = fmt.Println

func Test_PO_Request(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
}

func Test_PO_Response(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-1.xml"
	data := getPOXML(filename)

	var x ASIInventory

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)
}

func Test_PO_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Purchase-Response-Error-1.xml"
	data := getPOXML(filename)

	var x ASIInventory

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)
}

func Test_PurchaseOrder(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("TD_USER")
	client.Password = os.Getenv("TD_PASS")
	client.GetPriceAvail([]string{"209435"})
}

func getPOXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
