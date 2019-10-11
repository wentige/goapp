package controller

import (
	"html/template"
	"net/http"
)

type BookController struct{}

func (self BookController) RegisterRoute(g *echo.Group) {
	g.GET("/books", self.ReadList)
	g.GET("/book/:id", self.Detail)
}

func (BookController) ReadList(ctx echo.Context) error {
	return render(ctx, "books/list", data)
}

func (BookController) Detail(ctx echo.Context) error {
	return render(ctx, "books/detail", data)
}
