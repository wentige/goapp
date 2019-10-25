package debug

import (
	"encoding/json"
	"fmt"
	"strings"
)

//var pr = fmt.Println

// for reference only
func pr(v interface{}) {
	s := fmt.Sprintf("%#v\n", v)
	r := strings.NewReplacer("{", "{\n\t", "}", "\n}", ", ", ",\n\t")
	fmt.Println(r.Replace(s))
}

func dpr(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
