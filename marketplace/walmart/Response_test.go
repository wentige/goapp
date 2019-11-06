package walmart

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

var pr = fmt.Println
var ppr = pretty.Print

func Test_GetOrderInfoResponse(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)
}

func Test_Errors(t *testing.T) {
	assert := assert.New(t)

	filename := "fixtures/Errors.xml"
	bytes := getXML(filename)

	var x Errors
	err := xml.Unmarshal(bytes, &x)
	assert.Nil(err)

	assert.Equal(x.Error.Code, "INVALID_REQUEST_CONTENT.GMP_ORDER_API")
	assert.Equal(x.Error.Field, "data")
	assert.Equal(x.Error.Description, "Unable to process this request. The Line: 1 of PO: Y12613640 does not exist")
	assert.Equal(x.Error.Info, "Request content is invalid.")
	assert.Equal(x.Error.Severity, "ERROR")
	assert.Equal(x.Error.Category, "DATA")
	assert.Equal(x.Error.Causes, "")
	assert.Equal(x.Error.Identifiers, "")
}

func getXML(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}
