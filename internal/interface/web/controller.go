package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/dachichang/goath-stack-template/internal/interface/web/view/pages"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctl *Controller) Home(c *gin.Context) {
	render(c, http.StatusOK, pages.Home())
}
