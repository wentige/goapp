package debug

import (
	"fmt"
	"github.com/kr/pretty"
)

var active = true

// debug.Print(v)
func Print(vals ...interface{}) {
	if !active {
		return
	}
	for _, v := range vals {
		pretty.Print(v)
	}
}

// debug.Enable()
func Enable() {
	active = true
}

// debug.Disable()
func Disable() {
	active = false
}
