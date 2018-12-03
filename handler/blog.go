package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/huudung13/dungf1/models"
)

const (
	UserSessionKey = "UserSession"
)

func blogRouter(route chi.Router) {
	route.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			blogs, err := models.GetBlogs()
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			if err == nil {
				tmplHelper.Render(w, "blog_index", Map{"blogs": blogs, "user": user})
			}

		})
		r.Get("/view/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			BlogID := chi.URLParam(r, "id")
			blog, err := models.GetOnePost(strconv.Atoi(BlogID)) // Chuyển BlogID từ string thành int
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			if err == nil {
				tmplHelper.Render(w, "blog_view", Map{"blog": blog, "user": user})
			}
		})
		r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			tmplHelper.Render(w, "blog_create", Map{"user": user})
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
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			if err == nil {
				tmplHelper.Render(w, "blog_edit", Map{"blog": blog, "user": user})
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
		//-------------------Func for account----------------------------------------------

	})
}
