package view

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// PrependBaseURI prepends the base URI to the string
func PrependBaseURI(s string) string {
	return Config.BaseURI + s
}

// AssetTimePath returns a URL with the proper base uri and timestamp appended.
// Works for CSS and JS assets
func AssetTimePath(s string) (string, error) {
	// Determines if local or on the web
	if strings.HasPrefix(s, "//") {
		return s, nil
	}

	//s = strings.TrimLeft(s, "/")
	s = Config.AssetPath + strings.TrimLeft(s, "/")
	abs, err := filepath.Abs(s)
	if err != nil {
		return "", err
	}

	time, err2 := FileTime(abs)
	if err2 != nil {
		return "", err2
	}

	return PrependBaseURI(s + "?" + time), nil
}

// FileTime returns the modification time of the file
func FileTime(name string) (string, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return "", err
	}
	mtime := fi.ModTime().Unix()
	return fmt.Sprintf("%v", mtime), nil
}

// Validate returns true if all the required form values are passed
func Validate(req *http.Request, required []string) (bool, string) {
	for _, v := range required {
		if req.FormValue(v) == "" {
			return false, v
		}
	}

	return true, ""
}

// Repopulate updates the dst map so the form fields can be refilled
// useage: Refill any form fields
//	 view.Repopulate([]string{"email", "fullname", "message"}, r.Form, v.Vars)
func Repopulate(list []string, src url.Values, dst map[string]interface{}) {
	for _, v := range list {
		dst[v] = src.Get(v)
	}
}
