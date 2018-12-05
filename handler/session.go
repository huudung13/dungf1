package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	sessions "github.com/kataras/go-sessions"
)

var (
	sess = sessions.New(sessions.Config{
		Cookie:                      "mycookiesessionnameid",
		Expires:                     time.Hour * 2,
		DisableSubdomainPersistence: false,
	})
)

func sessionRoute(route chi.Router) {
	route.Route("/", func(r chi.Router) {
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			sess.Start(w, r).Clear()
			http.Redirect(w, r, "/blog", 302)
		})

	})
}
