package asi

import (
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
	assert.Nil(nil)
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
