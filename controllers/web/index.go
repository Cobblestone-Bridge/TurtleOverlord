package web

import (
	"net/http"

	"html/template"

	"github.com/Cobblestone-Bridge/TurtleOverlord/helpers"
	"github.com/zenazn/goji/web"
)

// Home page route
func (controller *Controller) Index(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)

	widgets := helpers.Parse(t, "index", nil)

	// With that kind of flags template can "figure out" what route is being rendered
	c.Env["IsIndex"] = true

	c.Env["Title"] = "TurtleOverlord"
	c.Env["Content"] = template.HTML(widgets)

	return helpers.Parse(t, "main", c.Env), http.StatusOK
}
