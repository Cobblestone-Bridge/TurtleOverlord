package system

import (
	"bytes"
	"html/template"

	"github.com/gorilla/sessions"
	"github.com/zenazn/goji/web"
	mgo "gopkg.in/mgo.v2"
)

type Controller struct {
}

func (controller *Controller) GetSession(c web.C) *sessions.Session {
	return c.Env["Session"].(*sessions.Session)
}

func (controller *Controller) GetTemplate(c web.C) *template.Template {
	return c.Env["Template"].(*template.Template)
}

func (controller *Controller) GetConfiguration(c web.C) *Configuration {
	return c.Env["Configuration"].(*Configuration)
}

func (controller *Controller) GetDatabase(c web.C) *mgo.Database {
	dbSession := c.Env["DBSession"].(*mgo.Session)
	return dbSession.DB(c.Env["DBName"].(string))
}

func (controller *Controller) Parse(t *template.Template, name string, data interface{}) string {
	var doc bytes.Buffer
	t.ExecuteTemplate(&doc, name, data)
	return doc.String()
}
