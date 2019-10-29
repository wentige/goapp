package dh

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pr = fmt.Println

func Test_XmlRequest(t *testing.T) {
	assert := assert.New(t)

	var x XmlRequest

	x.Request = "price-availability"
	x.Login.UserId = "USERID"
	x.Login.Password = "PASSWORD"
	x.Partnums = []string{"SKU-1", "SKU-2"}

	output, err := xml.MarshalIndent(x, "", "  ")
	assert.Nil(err)

	//fmt.Println(string(output))

	var z XmlRequest

	err = xml.Unmarshal(output, &z)
	assert.Nil(err)

	assert.Equal(x.Request, z.Request)
	assert.Equal(x.Login.UserId, z.Login.UserId)
	assert.Equal(x.Login.Password, z.Login.Password)
	assert.Equal(len(x.Partnums), 2)
	assert.Equal(x.Partnums, z.Partnums)
}

func Test_XmlResponse(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/PriceAvail-Response-1.xml"
	data := getXML(filename)

	var x XmlResponse

	err := xml.Unmarshal(data, &x)
	assert.Nil(err)

	assert.Equal(x.Status, "success")

	assert.Equal(len(x.Items), 2)
	assert.Equal(x.Items[0].Message, "")
	assert.Equal(x.Items[0].PartNum, "STEB8000100CA")

	assert.Equal(len(x.Items[0].BranchQty), 2)

	assert.Equal(x.Items[0].BranchQty[0].Branch, "Vancouver")
	assert.Equal(x.Items[0].BranchQty[0].Qty, "160")

	assert.Equal(x.Items[0].BranchQty[1].Branch, "Toronto")
	assert.Equal(x.Items[0].BranchQty[1].Qty, "483")

	assert.Equal(x.Items[1].PartNum, "MCUPMP")
	assert.Equal(x.Items[1].Message, "Model: MCUPMP has been discontinued and has No stock")
}

func Test_Errors(t *testing.T) {
}

func Test_Client(t *testing.T) {
	client := &Client{}
	client.Username = os.Getenv("DH_USER")
	client.Password = os.Getenv("DH_PASS")
	client.GetPriceAvail([]string{"DH-980001203CDN", "DH-XS708T100NESCA"})
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
