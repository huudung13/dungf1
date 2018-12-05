package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/huudung13/dungf1/helper"
	"github.com/huudung13/dungf1/models"
)

type (
	handlerHelper struct{}
	Map           map[string]interface{}
)

var (
	tmplHelper  *helper.TmplHelper
	formDecoder *form.Decoder
)

func init() {
	tmplHelper, _ = helper.NewTPL(helper.TmplConfig{"views", "html", "error_notFound"}, funcMap())
	formDecoder = form.NewDecoder()
}

func InitHandler(route *chi.Mux) {
	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHelper.Render(w, "home", Map{"user": models.Account{}})
	})
	route.Post("/upload/awata", uploadFile)
	route.Route("/blog", blogRouter)
	route.Route("/session", sessionRoute)
	route.Route("/auth", authRouter)

}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"json": func(data interface{}) string {
			_d, _ := json.Marshal(data)
			return string(_d)
		},
	}
}

func (h handlerHelper) Render(w http.ResponseWriter, tmplName string, data interface{}) {
	tmplHelper.Render(w, tmplName, data)
}

func (h handlerHelper) DecodeForm(r *http.Request, data interface{}) error {
	r.ParseForm()
	return formDecoder.Decode(data, r.Form)
}

func (h handlerHelper) DecodeMultipart(r *http.Request, data interface{}) error {
	r.ParseMultipartForm(1024 * 1024 * 2)
	return formDecoder.Decode(data, r.Form)
}
