package debug

import (
	"fmt"
)

var active = true

// debug.Print(v)
func Print(vals ...interface{}) {
	if !active {
		return
	}
	for _, v := range vals {
		fmt.Printf("%#v\n", v)
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
