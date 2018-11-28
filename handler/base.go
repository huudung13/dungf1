package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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

		tmplHelper.Render(w, "home", Map{"user": models.Blog{}})
	})
	route.Route("/blog", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			blogs, err := models.GetBlogs()
			fmt.Println(blogs, err)
			tmplHelper.Render(w, "blog_index", Map{"blogs": blogs})
		})
		r.Get("/view/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			BlogID := chi.URLParam(r, "id")
			blog, err := models.GetOnePost(strconv.Atoi(BlogID)) // Chuyển BlogID từ chuỗi thành số

			if err == nil {
				tmplHelper.Render(w, "blog_view", Map{"blog": blog})
			}

		})
		r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
			tmplHelper.Render(w, "blog_create", Map{"user": models.Blog{}})
		})
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
			// Code nhan gia tri tu form va luu tru vao database
			r.ParseForm()
			blog := models.Blog{
				Title:       r.FormValue("Title"),
				Description: r.FormValue("Description"),
				Content:     r.FormValue("Content"),
			}
			blog.Create()
			http.Redirect(w, r, "/blog", 302)
		})
		r.Get("/edit/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			BlogID := chi.URLParam(r, "id")
			blog, err := models.GetOnePost(strconv.Atoi(BlogID))
			if err == nil {
				tmplHelper.Render(w, "blog_edit", Map{"blog": blog})
			}

		})
		r.Post("/edit/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			// Code de xu ly khi sua 1 ban ghi du lieu
			BlogID := chi.URLParam(r, "id")
			ThisID, _ := strconv.Atoi(BlogID)
			r.ParseForm()
			blog := models.Blog{
				ID:          ThisID,
				Title:       r.FormValue("Title"),
				Description: r.FormValue("Description"),
				Content:     r.FormValue("Content"),
			}

			//fmt.Println(blog)
			blog.Update()
			http.Redirect(w, r, "/blog", 302)

		})
		r.Get("/delete/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			BlogID := chi.URLParam(r, "id")
			models.DeleteOnePost(strconv.Atoi(BlogID))
			http.Redirect(w, r, "/blog", 302)
		})
	})
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
