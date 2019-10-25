package view

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

var Config config
var TemplateFuncMap template.FuncMap
var TemplateCache = make(map[string]*template.Template)
var PartialTemplates []string

// The config can be changed outside view package
// view.Config.AssetPath = "asset/"
// view.Config.LayoutPath = "template/layout/"
// view.Config.DefaultLayout = "base"
// view.Config.PartialPath = "template/partial/"
// view.Config.PagePath = "template/page/"
// view.Config.TemplateExt = ".gotml"

type config struct {
	BaseURI       string
	AssetPath     string
	LayoutPath    string
	DefaultLayout string
	RootTemplate  string
	PartialPath   string
	PagePath      string
	TemplateExt   string
}

type View struct {
	Name    string
	Root    string
	Layout  string
	Caching bool
	Vars    map[string]interface{}
	request *http.Request
}

func init() {
	Config.BaseURI = "/"
	Config.AssetPath = "asset/"
	Config.LayoutPath = "template/layout/"
	Config.DefaultLayout = "base"
	Config.RootTemplate = "template/root"
	Config.PartialPath = "template/partial/"
	Config.PagePath = "template/page/"
	Config.TemplateExt = ".tmpl" // gotmpl

	TemplateFuncMap = GetTemplateFuncMap()

	// load partial templates
	PartialTemplates, _ = filepath.Glob(Config.PartialPath + "*" + Config.TemplateExt)
}

// v := view.New(r)
// v.Name = "about"
// v.Render(w)

func New(req *http.Request) *View {
	v := &View{}
	v.Layout = Config.DefaultLayout
	v.Root = Config.RootTemplate
	v.request = req
	v.Vars = make(map[string]interface{})
	v.Vars["UserId"] = 0
	v.Vars["UserRole"] = ""
	v.Vars["BaseURI"] = Config.BaseURI
	return v
}

func (v *View) Render(w io.Writer) {
	var filenames []string

	// check cache first
	if tmpl, ok := TemplateCache[v.Name]; ok {
		err := tmpl.ExecuteTemplate(w, v.Root, v.Vars)
		if err != nil {
			log.Println("Template Execute Error: ", err)
		}
		return
	}

	// partial template can be overwritten
	filenames = append(filenames, Config.LayoutPath+v.Layout+Config.TemplateExt)
	filenames = append(filenames, PartialTemplates...)
	filenames = append(filenames, Config.PagePath+v.Name+Config.TemplateExt)

	// parse templates
	tmpl, err := template.New(v.Name).Funcs(TemplateFuncMap).ParseFiles(filenames...)
	if err != nil {
		log.Println("Template Parse Error: ", err)
		return
	}

	// cache the templates
	TemplateCache[v.Name] = tmpl

	// execute templates
	err = tmpl.ExecuteTemplate(w, v.Root, v.Vars)
	if err != nil {
		log.Println("Template Execute Error: ", err)
	}
}
