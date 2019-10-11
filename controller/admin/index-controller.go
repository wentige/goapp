package admin

import (
	"html/template"
	"net/http"
)

type IndexController struct{}

func (self IndexController) RegisterRoute(g *echo.Group) {
	g.GET("/", self.Index)
}

func (IndexController) Index(w http.ResponseWriter, r *http.Request) error {
	return render(w, "index/index", data)
}
