package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PagesController struct{}

func NewPagesController() *PagesController {
	return &PagesController{}
}

func (pc *PagesController) Mount(g *echo.Group) {
	g.GET("/hello", pc.hello)
	// g.GET("/version", pc.version)
}

func (pc *PagesController) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
