package view

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"strings"
)

// GetTemplateFuncMap returns a map of functions that are usable in templates
func GetTemplateFuncMap() template.FuncMap {
	f := make(template.FuncMap)

	f["JS"] = func(s string) template.HTML {
		path, err := AssetTimePath(s)
		if err != nil {
			log.Println("JS Error:", err)
			return template.HTML("<!-- JS Error: " + s + " -->")
		}

		return template.HTML(`<script type="text/javascript" src="` + path + `"></script>`)
	}

	f["CSS"] = func(s string) template.HTML {
		path, err := AssetTimePath(s)
		if err != nil {
			log.Println("CSS Error:", err)
			return template.HTML("<!-- CSS Error: " + s + " -->")
		}

		return template.HTML(`<link rel="stylesheet" type="text/css" href="` + path + `" />`)
	}

	f["LINK"] = func(path, name string) template.HTML {
		return template.HTML(`<a href="` + PrependBaseURI(path) + `">` + name + `</a>`)
	}

	f["NOESCAPE"] = func(text string) template.HTML {
		return template.HTML(text)
	}

	// Same as NOESCAPE, more meaningful
	f["htmlsafe"] = func(text string) template.HTML {
		return template.HTML(text)
	}

	f["uppercase"] = strings.ToUpper
	f["lowercase"] = strings.ToLower
	f["title"] = strings.ToTitle

	f["RANDIMG"] = func() template.HTML {
		num := rand.Intn(11)
		return template.HTML(fmt.Sprintf("%v", num))
	}

	f["RANDIMGSLIDER"] = func() template.HTML {
		num := rand.Intn(4)
		return template.HTML(fmt.Sprintf("%v", num))
	}

	return f
}
